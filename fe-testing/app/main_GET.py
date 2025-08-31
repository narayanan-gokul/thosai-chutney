from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import HTMLResponse
from pathlib import Path

app = FastAPI()

# Allow frontend to fetch from backend
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Adjust for production
    allow_methods=["*"],
    allow_headers=["*"],
)

# ----- Serve HTML Pages -----
@app.get("/consumer_supply", response_class=HTMLResponse)
def serve_consumer_supply():
    path = Path("templates/consumer_supply.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/distributor_dashboard", response_class=HTMLResponse)
def serve_distributor_dashboard():
    path = Path("templates/distributor_dashboard.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/supplier_dashboard", response_class=HTMLResponse)
def serve_supplier_dashboard():
    path = Path("templates/supplier_dashboard.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/supplier_Welcome", response_class=HTMLResponse)
def serve_supplier_dashboard():
    path = Path("templates/supplier_Welcome.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/signup_distributor", response_class=HTMLResponse)
def serve_supplier_dashboard():
    path = Path("templates/signup_distributor.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/signup_supplier", response_class=HTMLResponse)
def serve_supplier_dashboard():
    path = Path("templates/signup_supplier.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

@app.get("/consumer_Welcome", response_class=HTMLResponse)
def serve_supplier_dashboard():
    path = Path("templates/consumer_Welcome.html")
    return HTMLResponse(content=path.read_text(), status_code=200)

# ----- API Endpoints -----

# Get list of food items with max quantity
@app.get("/api/items")
def get_items():
    return [
        {"id": 1, "name": "Rice (kg)", "max": 10},
        {"id": 2, "name": "Eggs (units)", "max": 12},
        {"id": 3, "name": "Milk (litres)", "max": 10},
        {"id": 4, "name": "Chicken", "max": 5},
        {"id": 5, "name": "Bread (packets)", "max": 3}
    ]

# Get consumer info
@app.get("/api/consumer_info")
def get_consumer_info():
    return {"first_name": "John", "last_name": "Doe", "postcode": 2000}

# Get all orders for a distributor (sum of consumer requests)
@app.get("/api/distributor_orders")
def get_distributor_orders():
    return {
        "total": {"Rice": 20, "Eggs": 35, "Milk": 15, "Chicken": 8, "Bread": 10},
        "individual": [
            {"consumer": "John Doe", "order": {"Rice": 5, "Eggs": 6}},
            {"consumer": "Jane Smith", "order": {"Milk": 5, "Bread": 3}},
        ]
    }

# Get supplier stock info
@app.get("/api/supplier_stock")
def get_supplier_stock():
    return {
        "Rice": 50,
        "Eggs": 100,
        "Milk": 40,
        "Chicken": 20,
        "Bread": 30
    }

@app.get("/api/distributor_info")
def distributor_info(dist_id: int):
    # Replace this with real DB query
    return {"name": "FoodHub A", "postcode": 2000}
