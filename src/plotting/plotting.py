import requests
import json
from typing import Callable

url = "http://localhost:8080/"

def create_function(coeffs: list[list[int]], signs: list[int]) -> Callable[float, float]: 
	def func(x: float) -> float:
		res = 0
		for i, monomial in enumerate(coeffs):
			if signs[i] == 43:
				res += monomial[0] * x ** monomial[1] 
			else:
				res -= monomial[0] * x ** monomial[1]
		return res
	return func

def main() -> None:
	while True:
		function = input("Enter the desired function: ")
		params = {
		    "function": function
		}

		response = requests.get(url, params=params)
		fn_result = response.json()
		func = create_function(fn_result["coeffs"], fn_result["signs"])
		print(func(1))
	
if __name__ == "__main__":
	main()
