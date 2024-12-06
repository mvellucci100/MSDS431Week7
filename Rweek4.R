# The Anscombe Quartet (R)

# Load necessary libraries
library(ggplot2)

# Define the Anscombe data frame
anscombe <- data.frame(
  x1 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x2 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x3 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
  x4 = c(8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8),
  y1 = c(8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68),
  y2 = c(9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74),
  y3 = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73),
  y4 = c(6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89)
)

# Measure execution time for each dataset
for (i in 1:4) {
  # Prepare the design matrix
  design_matrix <- cbind(1, anscombe[[paste0("x", i)]])  # Add intercept column
  
  # Fit the linear regression model
  start_time <- Sys.time()
  model <- lm(anscombe[[paste0("y", i)]] ~ design_matrix)
  end_time <- Sys.time()
  
  # Print the summary
  print(summary(model))
  
  # Calculate and print execution time
  execution_time <- end_time - start_time
  cat(sprintf("Execution time for Dataset %d: %.6f seconds\n\n", i, execution_time))
}

# Create scatter plots
par(mfrow=c(2, 2))  # Set up a 2x2 plotting area

plot(anscombe$x1, anscombe$y1, main='Set I', xlab='x1', ylab='y1', xlim=c(2, 20), ylim=c(2, 14))
plot(anscombe$x2, anscombe$y2, main='Set II', xlab='x2', ylab='y2', xlim=c(2, 20), ylim=c(2, 14))
plot(anscombe$x3, anscombe$y3, main='Set III', xlab='x3', ylab='y3', xlim=c(2, 20), ylim=c(2, 14))
plot(anscombe$x4, anscombe$y4, main='Set IV', xlab='x4', ylab='y4', xlim=c(2, 20), ylim=c(2, 14))

# Save the plots as a PDF
dev.copy(pdf, "fig_anscombe_R.pdf")
dev.off()

# Suggestions for the student:
# See if you can develop a quartet of your own, 
# or perhaps just a duet, two very different data sets 
# with the same fitted model.
