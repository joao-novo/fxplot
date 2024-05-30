import requests
import numpy as np
import plotly.express as px
import pandas as pd
from typing import Callable

url = "http://localhost:8080/"


def create_function(
    coeffs: list[list[int]], signs: list[int]
) -> Callable[[float], float]:
    def func(x: float) -> float:
        res = 0
        for i, monomial in enumerate(coeffs):
            match signs[i]:
                case 43:  # 43 == '+'
                    res += monomial[0] * x ** monomial[1]
                case 45:  # 45 == '-'
                    res -= monomial[0] * x ** monomial[1]
                case _:
                    continue
        return res

    return func


def generate_values(func: Callable[[float], float]) -> pd.DataFrame:
    vals = np.linspace(-100, 100, 1000000)
    return pd.DataFrame([func(y) for y in vals])


def main() -> None:
    function = input("Enter the desired function: ")
    params = {"function": function}

    response = requests.get(url, params=params)
    fn_result = response.json()
    func = create_function(fn_result["coeffs"], fn_result["signs"])
    df = generate_values(func)
    fig = px.line(df)
    fig.write_html("../../fig.html")


if __name__ == "__main__":
    main()
