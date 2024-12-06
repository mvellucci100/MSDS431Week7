import numpy as np
import statsmodels.api as sm
import timeit

# Define the datasets
datasets = {
    "Dataset 1": np.array([[10, 8.04], [8, 6.95], [13, 7.58], [9, 8.81], [11, 8.33],
                            [14, 9.96], [6, 7.24], [4, 4.26], [12, 10.84], [7, 4.82], [5, 5.68]]),
    "Dataset 2": np.array([[10, 9.14], [8, 8.14], [13, 8.74], [9, 8.77], [11, 9.26],
                            [14, 8.10], [6, 6.13], [4, 3.10], [12, 9.13], [7, 7.26], [5, 4.74]]),
    "Dataset 3": np.array([[10, 7.46], [8, 6.77], [13, 12.74], [9, 7.11], [11, 7.81],
                            [14, 8.84], [6, 6.08], [4, 5.39], [12, 8.15], [7, 6.42], [5, 5.73]]),
    "Dataset 4": np.array([[8, 6.58], [8, 5.76], [8, 7.71], [8, 8.84], [8, 8.47],
                            [8, 7.04], [8, 5.25], [19, 12.50], [8, 5.56], [8, 7.91], [8, 6.89]])
}

# Perform linear regression and print results
for name, data in datasets.items():
    x = sm.add_constant(data[:, 0])  # Add a constant term for the intercept
    y = data[:, 1]
    model = sm.OLS(y, x).fit()
    print(f"{name} - Intercept: {model.params[0]:.2f}, Slope: {model.params[1]:.2f}")

# Benchmarking the linear regression
def benchmark():
    for data in datasets.values():
        x = sm.add_constant(data[:, 0])
        y = data[:, 1]
        model = sm.OLS(y, x).fit()

# Measure execution time
execution_time = timeit.timeit(benchmark, number=1000)
print(f"Average execution time over 1000 runs: {execution_time:.4f} seconds")
