# Food catalog with limits
food_items = {
    1: {"name": "Rice", "max_quantity": 10},
    2: {"name": "Wheat", "max_quantity": 8},
    3: {"name": "Milk", "max_quantity": 5},
    4: {"name": "Eggs", "max_quantity": 12},
    5: {"name": "Bread", "max_quantity": 6}
}

def process_order(request_data: dict) -> dict:
    """
    Process a food order and return only valid items.
    Invalid items (wrong ID or exceeding limits) are ignored.

    Input format:
    {
        "items": [
            {"food_id": 1, "quantity": 3},
            {"food_id": 3, "quantity": 7}
        ]
    }

    Output format:
    {
        "order": [
            {"food_id": 1, "name": "Rice", "quantity": 3}
        ]
    }
    """

    food_items = {
    1: {"name": "Rice", "max_quantity": 10},
    2: {"name": "Wheat", "max_quantity": 8},
    3: {"name": "Milk", "max_quantity": 5},
    4: {"name": "Eggs", "max_quantity": 12},
    5: {"name": "Bread", "max_quantity": 6}
}

    if not request_data or "items" not in request_data:
        return {"order": []}

    consumer_order = []

    for item in request_data["items"]:
        food_id = item.get("food_id")
        quantity = item.get("quantity")

        # Only accept if valid food ID and within allowed range
        if food_id in food_items:
            max_quantity = food_items[food_id]["max_quantity"]
            if isinstance(quantity, int) and 0 < quantity <= max_quantity:
                consumer_order.append({
                    "food_id": food_id,
                    "name": food_items[food_id]["name"],
                    "quantity": quantity
                })

    return {"order": consumer_order}

