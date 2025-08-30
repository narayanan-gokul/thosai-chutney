def login("users_database.txt"):
    print("Welcome back to Fighting Food Fragments!")
    logged_in = False
    while logged_in == False:
        try:
            input_username = input("Please enter your user name: ")
            input_password = input("Please enter your password: ")
            with open("users_database.txt", 'r') as file:
                content = file.read()
                if input_username in content and input_password in content:
                    logged_in = True
                else:
                    print("User name or password invalid, please check and try again.")
        except FileNotFoundError:
            print(f"Error: Users database file not found")
    return True

