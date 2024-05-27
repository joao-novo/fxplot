import requests
import json

url = "http://localhost:8080/"

def main() -> None:
    function = input("Enter the desired function: ")
    params = {
        "function": function
    }
    
    response = requests.get(url, params=params)
    fn_result = response.json()
    print(fn_result)

if __name__ == "__main__":
    main()
