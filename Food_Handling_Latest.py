# ============================
# Food Catalog
# ============================
food_catalog = {
    1: {"name": "Rice", "max_quantity": 10},
    2: {"name": "Wheat", "max_quantity": 8},
    3: {"name": "Milk", "max_quantity": 5},
    4: {"name": "Eggs", "max_quantity": 12},
    5: {"name": "Bread", "max_quantity": 6}
}

# ============================
# 1) Consumer Side
# ============================
def process_consumer_request(request_data: dict, postcode: str) -> dict:
    """
    Validate consumer request and attach the chosen foodhub.
    Each consumer can choose only one foodhub in their postcode.
    """
    foodhub = request_data.get("foodhub")
    if not foodhub:
        raise ValueError("Consumer must choose a foodhub.")

    cleaned_items = []

    for item in request_data.get("items", []):
        fid = item.get("food_id")
        qty = item.get("quantity")

        if fid in food_catalog and isinstance(qty, int) and 0 < qty <= food_catalog[fid]["max_quantity"]:
            cleaned_items.append({
                "food_id": fid,
                "name": food_catalog[fid]["name"],
                "quantity": qty
            })

    return {
        "postcode": postcode,
        "foodhub": foodhub,
        "items": cleaned_items
    }

# ============================
# 2) FoodHub / Distributor Side
# ============================
def aggregate_foodhub_orders(consumer_orders: list) -> dict:
    """
    Sum all requests from consumers for each foodhub.
    """
    foodhub_totals = {}

    for order in consumer_orders:
        hub = order["foodhub"]
        if hub not in foodhub_totals:
            foodhub_totals[hub] = {}

        for item in order["items"]:
            name = item["name"]
            qty = item["quantity"]
            foodhub_totals[hub][name] = foodhub_totals[hub].get(name, 0) + qty

    return foodhub_totals

# ============================
# 3) Supplier Side
# ============================
def supplier_inventory(supplier_name: str, inventory: list) -> dict:
    """
    Store supplier's inventory for the week.
    Inventory example: [{"food_id": 1, "quantity": 100}, {"food_id": 2, "quantity": 50}]
    """
    inventory_dict = {}
    for item in inventory:
        fid = item.get("food_id")
        qty = item.get("quantity", 0)
        if fid in food_catalog and qty >= 0:
            inventory_dict[food_catalog[fid]["name"]] = qty

    return {
        "supplier": supplier_name,
        "inventory": inventory_dict
    }

# ============================
# Example Usage (can be called from frontend)
# ============================
if __name__ == "__main__":
    # Consumer requests
    consumer1 = process_consumer_request(
        {"items": [{"food_id": 1, "quantity": 3}, {"food_id": 3, "quantity": 2}], "foodhub": "FoodHub_A"},
        postcode="2000"
    )
    consumer2 = process_consumer_request(
        {"items": [{"food_id": 1, "quantity": 4}, {"food_id": 5, "quantity": 2}], "foodhub": "FoodHub_A"},
        postcode="2000"
    )
    consumer3 = process_consumer_request(
        {"items": [{"food_id": 2, "quantity": 5}], "foodhub": "FoodHub_B"},
        postcode="2000"
    )

    consumer_orders = [consumer1, consumer2, consumer3]

    # FoodHub aggregation
    hub_totals = aggregate_foodhub_orders(consumer_orders)
    print("FoodHub Totals:", hub_totals)
    # Example Output: {'FoodHub_A': {'Rice': 7, 'Milk': 2, 'Bread': 2}, 'FoodHub_B': {'Wheat': 5}}

    # Supplier inventory
    coles_inventory = supplier_inventory("Coles", [{"food_id": 1, "quantity": 100}, {"food_id": 2, "quantity": 50}, {"food_id": 3, "quantity": 30}])
    woolworths_inventory = supplier_inventory("Woolworths", [{"food_id": 4, "quantity": 60}, {"food_id": 5, "quantity": 40}])

    print("Supplier Inventories:", coles_inventory, woolworths_inventory)
