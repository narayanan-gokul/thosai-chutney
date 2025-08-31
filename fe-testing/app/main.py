from fastapi import FastAPI, HTTPException
from fastapi.responses import HTMLResponse, JSONResponse
from pydantic import BaseModel
from typing import Optional
from fastapi.middleware.cors import CORSMiddleware
from typing import List, Dict
app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

# -------------------------------
# MODELS
# -------------------------------
class ConsumerSignup(BaseModel):
    first_name: str
    last_name: str
    password: str
    postcode: str

class DistributorSignup(BaseModel):
    name: str
    password: str
    postcode: str

class SupplierSignup(BaseModel):
    name: str
    password: str
    postcode: str

class LoginRequest(BaseModel):
    role: str          # "consumer", "supplier", "distributor"
    email: str
    password: str

# Example data model for an order
class SupplierOrder(BaseModel):
    consumer_name: str
    item: str
    quantity: int

class ConsumerWelcomeRequest(BaseModel):
    userId: str
    postcode: str

class ConsumerSupplyRequest(BaseModel):
    userId: str
    foodhubId: int
    order: Dict[int, int]  # {item_id: quantity}

class LoginRequest(BaseModel):
    userId: str

dummy_users = {
    "consumer": [{"email":"consumer@test.com","password":"123","name":"Alice","user_id":1}],
    "supplier": [{"email":"supplier@test.com","password":"123","name":"Bob","user_id":2}],
    "distributor": [{"email":"distributor@test.com","password":"123","name":"Charlie","user_id":3}],
}

foodhubs_data = {
    "2000": [{"id": 1, "name": "Sydney Central Hub"}, {"id": 2, "name": "Parramatta Hub"}],
    "2016": [{"id": 3, "name": "Newtown Hub"}]
}

FOODHUBS = {
    "2000": [{"id": 1, "name": "Sydney Central Hub"}, {"id": 2, "name": "Parramatta Hub"}],
    "2016": [{"id": 3, "name": "Newtown Hub"}]
}

SUPPLIES = [
    {"id": 1, "name": "Rice (in kg)", "max": 10},
    {"id": 2, "name": "Eggs (per unit)", "max": 12},
    {"id": 3, "name": "Milk (in litres)", "max": 10},
    {"id": 4, "name": "Chicken", "max": 5},
    {"id": 5, "name": "Bread (per packet)", "max": 3}
]

# Dummy data for testing
dummy_supplier_orders = [
    {"consumer_name": "Alice", "item": "Rice", "quantity": 5},
    {"consumer_name": "Bob", "item": "Eggs", "quantity": 12},
    {"consumer_name": "Charlie", "item": "Milk", "quantity": 3},
    {"consumer_name": "Alice", "item": "Milk", "quantity": 1}
]

dummy_distributor_orders = [
    {"consumer_name": "Alice", "item": "Rice", "quantity": 5},
    {"consumer_name": "Bob", "item": "Eggs", "quantity": 12},
    {"consumer_name": "Charlie", "item": "Milk", "quantity": 3},
    {"consumer_name": "Alice", "item": "Milk", "quantity": 1}
]



# Dummy DB for testing (would come from signup)
users_db = {
    "consumer-123": {"role": "Consumer"},
    "distributor-456": {"role": "Distributor"},
}
# -------------------------------
# ROUTES
# -------------------------------
@app.post("/api/consumer/signup")
def signup_consumer(consumer: ConsumerSignup):
    print("✅ New consumer:", consumer.dict())
    return {"message": "Consumer signup successful", "data": consumer.dict()}

@app.post("/api/distributor/signup")
def signup_distributor(distributor: DistributorSignup):
    print("✅ New distributor:", distributor.dict())
    return {"message": "Distributor signup successful", "data": distributor.dict()}

@app.post("/api/supplier/signup")
def signup_supplier(supplier: SupplierSignup):
    print("✅ New supplier:", supplier.dict())
    return {"message": "Supplier signup successful", "data": supplier.dict()}

@app.get("/api/consumer/foodhubs")
def get_foodhubs(postcode: int):
    # Return FoodHubs in the consumer's area
    available_hubs = [hub for hub in foodhubs_data if abs(hub["postcode"] - postcode) == 0]
    return {"foodhubs": available_hubs}
    
@app.get("/consumer_welcome", response_class=HTMLResponse)
def serve_consumer_welcome():
    html_path = "templates/consumer_welcome.html"
    html_content = open(html_path, "r").read()
    return HTMLResponse(content=html_content)


@app.post("/api/login")
def login(request: LoginRequest):
    role = request.role.lower()
    email = request.email
    password = request.password

    if role not in dummy_users:
        raise HTTPException(status_code=400, detail="Invalid role")

    user_list = dummy_users[role]
    user = next((u for u in user_list if u["email"] == email and u["password"] == password), None)

    if not user:
        raise HTTPException(status_code=401, detail="Invalid email or password")

    redirect_map = {
        "consumer": "/consumer_dashboard",
        "supplier": "/supplier_dashboard",
        "distributor": "/distributor_dashboard"
    }

    return {
        "message": "Login successful",
        "access_token": "dummy_token_123",  # in real app we generate JWT
        "role": role,
        "user_id": user["user_id"],
        "name": user["name"],
        "redirect_url": redirect_map[role]
    }

@app.get("/api/supplier_orders")
def get_supplier_orders():
    return JSONResponse(content=dummy_supplier_orders)
@app.get("/api/distributor_orders")
def get_distributor_orders():
    return JSONResponse(content=dummy_distributor_orders)

@app.post("/api/login/consumer")
def login_consumer(data: LoginRequest):
    user = users_db.get(data.userId)
    if user and user["role"] == "Consumer":
        return {
            "message": "Login successful",
            "role": "Consumer",
            "userId": data.userId
        }
    raise HTTPException(status_code=401, detail="Invalid Consumer userId")


@app.post("/api/login/distributor")
def login_distributor(data: LoginRequest):
    user = users_db.get(data.userId)
    if user and user["role"] == "Distributor":
        return {
            "message": "Login successful",
            "role": "Distributor",
            "userId": data.userId
        }
    raise HTTPException(status_code=401, detail="Invalid Distributor userId")

@app.post("/api/welcome/consumer")
def consumer_welcome(req: ConsumerWelcomeRequest):
    hubs = FOODHUBS.get(req.postcode, [])
    return {"message": "Foodhubs fetched successfully", "foodhubs": hubs}


@app.post("/api/supply/consumer")
def consumer_supply(req: ConsumerSupplyRequest):
    # Save order in DB (mock here)
    return {
        "message": "Order submitted successfully",
        "userId": req.userId,
        "foodhubId": req.foodhubId,
        "order": req.order
    }


@app.get("/api/items/consumer")
def get_items():
    return {"items": SUPPLIES}