def check_supply(supplier_inventory, hub_totals):
    for food_id in supplier_inventory:
        for food_id in hub_totals:
            if supplier_inventory.food_id.quantity >= hub_totals.food_id.quantity:
                supplier_inventory.food_id.quantity -= hub_totals.food_id.quantity
            else:
                return Unstockable
            return Stockable
            
#Needs if statement to check that food hub has recieved all the food needed to fufill requests
def fufillable(hub_totals, consumer):
    for food_id in hub_totals:
        for food_id in consumer:
            if consumer.food_id.quantity <= hub_totals.food_id.quantity:
                hub_totals.food_id.quantity -= consumer.food_id.quantity
            else:
                return Unstockable
            return Stockable
