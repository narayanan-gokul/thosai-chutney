CREATE TRIGGER validate_order_quantity
BEFORE INSERT OR UPDATE ON cart
FOR EACH ROW
	EXECUTE FUNCTION check_order_less_than_max();
