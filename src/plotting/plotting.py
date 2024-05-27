import requests
import json
import numpy as np
import numpy.typing as npt
from typing import Callable

url = "http://localhost:8080/"

def create_function(coeffs: list[list[int]], signs: list[int]) -> Callable[float, float]: 
	def func(x: float) -> float:
		res = 0
		for i, monomial in enumerate(coeffs):
			match signs[i]:
				case 43: # 43 == '+'
					res += monomial[0] * x ** monomial[1] 
				case 45: # 45 == '-'
					res -= monomial[0] * x ** monomial[1]
				case _:
					continue
		return res
	return func

def generate_values(func: Callable[float, float]) -> npt.NDArray[float]:
	vals = np.linspace(-100, 100, 1000000)
	return np.array([y for y in func(vals)])



def main() -> None:
	while True:
		function = input("Enter the desired function: ")
		params = {
			"function": function
		}

		response = requests.get(url, params=params)
		fn_result = response.json()
		func = create_function(fn_result["coeffs"], fn_result["signs"])

	
if __name__ == "__main__":
	main()
