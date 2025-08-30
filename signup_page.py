def signup("user_data.txt"):
    print("Welcome to Fighting Food Fragments!")
    postcode_valid = false
    type_valid = false
    
    while type_valid = false:
        try:
            typeuser = int(input(
                """
                What are your needs? Press:
                1 - I am individual in need of food.
                2 - I am a food supplier that has food to give away.
                3 - I am a food hub that needs more food.
                """
            ))
            type_valid = true
        except TypeError:
            print("Please try again (1-3)")
            
    username = input("Please enter your name: ")
    
    while postcode_valid = false:
        try:
            postcode = int(input("Where are you? Please enter your postcode: "))
            if postcode >2000 and len(postcode)== 4:
                postcode_valid = true
            else:
                print("Please enter a valid postcode.")
        except typeError:
            print("Please enter a valid postcode.")
    
    password = input("Please enter your password: ")
    print(f"Your user name will be {username} and your password is {password}.")
    with open("user_data.txt", "a") as f:
            f.write("typeuser, username, postcode, password")
    return True

