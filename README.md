# MSDS431Week7
# Difference Between Automated Programming, AI-Assisted Programming, and AI-Generated Code

The terms **automated programming**, **AI-assisted programming**, and **AI-generated code** represent different levels and types of AI involvement in software development. Below is an overview of their differences.

---

## 1. Automated Programming
**Definition**: Automated programming refers to the use of tools or systems to handle repetitive and mechanical aspects of coding, often based on predefined templates or rules. This can include code generation, error-checking, formatting, testing, or deployment processes that do not necessarily involve AI but are more driven by automation scripts or pre-programmed algorithms.

## Automated Code Generation for Go Code

To apply automated code generation for the week4main.go file, we can follow a few strategies to automate or simplify repetitive parts of the program. In this case, we would focus on automating the creation of datasets, performing linear regression for multiple datasets, and handling common tasks like error checking.

### 1. Automate Dataset Generation
Since the datasets (`c1`, `c2`, `c3`, `c4`) follow a repetitive pattern of struct creation, we could automate the creation of these datasets by writing a script or using a tool that generates these coordinates based on a simple configuration or input.
For example, we could automate the generation of similar datasets using a configuration file and a Go code generator.

### 2. Automate Repetitive Linear Regression Calls
The process of calling `stats.LinearRegression()` on multiple datasets is repetitive. We can automate this by creating a function that takes any dataset and performs the regression. This will remove redundancy from the code and will refactor the code.

### 3. Automate Error Handling and Output
Rather than manually writing the error handling and output for each dataset, we could write a reusable function to handle this, and then loop through all datasets programmatically.

See AutomatedCodeGeneration.go file for sample output.
---

## 2. AI-Assisted Programming (using Github Copilot)
**Definition**: AI-assisted programming involves tools that leverage machine learning (ML) or AI models to enhance the productivity of developers. These tools **suggest** improvements, detect errors, provide code completions, or recommend solutions based on patterns identified in the developer's code or previous examples. The developer remains in control and makes the final decision.

### Suggestions for Statistical Calculations:
 GitHub Copilot may suggest common Go patterns or functions. For example, when calculating regression or preparing datasets, Copilot can recommend code snippets or optimization suggestions.

### Optimizations:
GitHub Copilot can assist by suggesting code refactoring or performance optimizations, such as:
- More efficient ways to loop through datasets.
- Utilizing Go's concurrency features.
- Handling errors more effectively.

### Code Example for Dataset Preparation:
For preparing datasets, Copilot might suggest using Go's built-in slice and map functions, helping streamline dataset processing.

Here’s an example of what Copilot might suggest for error handling in code:

```go
if err := stats.LinearRegression(c1); err != nil {
    log.Fatalf("Error in Dataset 1: %v", err)
}
```


## 3. AI-Generated Code (Using Chat GPT)
**Definition**: AI-generated code refers to the **complete generation of code** by an AI system with little or no human input. Here, AI models like large language models (LLMs) can take a developer’s request (a problem statement, requirements, or prompt) and generate full or partial code automatically, often based on previous patterns and training data.

For this section I used Chat GPT to generate Go code and assist with the task. The free verison of Chat GPT is accessible at  [https://chat.openai.com](https://chat.openai.com).

## Prompting ChatGPT
We will prompt ChatGPT to generate Go code for linear regression calculations. Below are some examples of how we can phrase our prompts:

- **Prompt1**: "Write Go code to perform linear regression on a dataset of x and y coordinates using the montanaflynn/stats package."
- **Prompt2**: "How would I modify the existing Go code to handle errors in a more structured way for each dataset?"
- **Prompt3**: "Generate Go code to calculate the linear regression of two variables using the montanaflynn/stats package."
- Outputs live in the repo labeled ChatGPT_Prompt1_output, ChatGPT_Prompt2_output, ChatGPT_Prompt3_output

## Summary
In my experience, Automated Code generation was very useful for non-intelligent, routine tasks such as code formatting, compiling, and basic error checking. However, it does not contribute significantly to solving higher-level programming challenges, which require a human touch. It can save significant time by eliminating the need for manual formatting and other mechanical tasks, allowing developers to focus on more complex problem-solving. However, its usefulness is limited to repetitive, well-defined tasks. For the startup, automated programming could be implemented to handle tasks like build automation, code formatting, and testing. However, it will not replace the need for developers in more complex, creative work like developing product features or managing business logic.

Using GitHub Copilot significantly enhanced the development experience by offering real-time suggestions for common coding patterns, optimizations, and even error handling. For example, Copilot suggested optimized loops and better ways to handle errors, such as using `log.Fatalf` instead of `fmt.Println`. This was particularly useful for a junior developer or someone who is still becoming familiar with Go. AI-assisted programming helps junior developers write better code faster by suggesting best practices and reducing the cognitive load involved in writing boilerplate code. This allows developers to focus more on problem-solving and product development while also reducing the amount of time spent on routine coding tasks. For the startup, AI-assisted programming tools like GitHub Copilot can be extremely beneficial in increasing developer productivity, especially for junior developers. It allows for faster onboarding and helps reduce errors and inefficiencies. I would recommend that the startup adopt tools like GitHub Copilot for its development team.

Using ChatGPT to generate Go code for linear regression calculations demonstrated the potential of AI to autonomously produce useful code snippets based on simple prompts. In this case, ChatGPT not only generated complete Go code for linear regression but also provided suggestions for more structured error handling. The process of interacting with ChatGPT was highly efficient, though some fine-tuning was necessary to align the code with specific project requirements.AI-generated code can greatly reduce the time needed to write boilerplate code or solve well-defined problems, making it useful for prototyping, rapid development, or exploring unfamiliar programming languages. However, human oversight is still necessary to ensure the generated code is correct, secure, and suitable for production use. For the startup, AI-generated code is an excellent tool for speeding up development, especially in the initial phases of product creation. It can be used for prototyping and generating boilerplate code quickly, but developers should review and test the generated code carefully. Given the potential for time-saving, I would recommend leveraging AI-generated code for certain aspects of development but maintaining a team of skilled developers to handle more complex tasks and ensure code quality.


---

