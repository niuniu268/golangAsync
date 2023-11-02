# Waitgroup and Async

## Advantage of waitgroup
- When the program finishes a register, the waitgroup should add 1. WaitGroup.add(n) means to launch N goroutine
- When the program finishes a report, the waitgroup should minus 1. WaitGroup.Done and one goroutine exits
- Waitgroup.wait aims to wait all goroutines finish. During waiting, the main goroutine is congest until all goroutine finish

![sync](img/Screenshot 2023-11-02 at 09.31.59.png)
![async](img/Screenshot 2023-11-02 at 09.32.15.png

## Waitgroup and Channel in order to aggregate all data

### Version 1
![channel1](img/Screenshot 2023-11-02 at 09.33.38.png)

### Version 2
![channel2](img/Screenshot 2023-11-02 at 09.33.55.png)

### Version 3
![channel3](img/Screenshot 2023-11-02 at 09.34.25.png)

## Explanation

### WaitGroup.add
![explanation1](img/Screenshot 2023-11-02 at 09.35.21.png)

### WaitGroup.wait
![explanation2](img/Screenshot 2023-11-02 at 09.41.46.png)


