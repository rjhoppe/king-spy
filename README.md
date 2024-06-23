# Overview
King SPY is a CLI application written in Go with the Cobra library that provides a suite of stock analysis tools to help investors discover outperforming equities within a given time period.

The baseline of "outperformance" is defined by outperforming the S&P 500 during a specified period of time, which is for convenience represented by the popular SPDR S&P 500 ETF (SPY).

Many of the tools provided in this application are not offered elsewhere (that I could find) and were included to streamline investor analysis.

# Example Output

![image](https://github.com/rjhoppe/king-spy/assets/48058874/eb0c3341-27af-40cd-9fa3-986813268cd7)

# Command List
Here is a list of all the commands you can run using King SPY:
* king-spy
* c2s (compare to SPY)
* c2t (compare to ticker)
* chart
* high
* low
* news
* random
* sectors
* all

# Description of Commands
### king-spy ###
Returns info about the application as well as descriptions of each command and usage examples
```
king-spy
```

### c2s ###
Compares a ticker's performance to the SP500 over a specified time period
```
king-spy c2s [ticker]
```
c2s with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy c2s [ticker] -t=[time period]
```
NOTE: When passing a -t flag, the ticker(s) must have existed on the NYSE for the entire duration of the time period or else you will receive an error

### c2t ###
Compares one ticker's performance to another ticker over a specified time period
```
king-spy c2t [ticker1] [ticker2]
```
c2t with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy c2t [ticker1] [ticker2]  -t=[time period]
```

### chart ###
Opens a stock chart for a specified entity in your default browser. Charting is provided by StockCharts.com
```
king-spy chart [ticker]
```

### high ###
Returns a ticker's percentage and dollar decrease from a recent high
```
king-spy high [ticker]
```
high with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy high [ticker] -t=[time period]
```

### low ###
Returns a ticker's percentage and dollar increase from a recent low
```
king-spy low [ticker]
```
low with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy low [ticker] -t=[time period]
```

### news ###
Returns the 5 most recent news headlines for a supplied ticker (headlines provided by Benzinga)
```
king-spy news [ticker]
```

### random ###
Compares the performance of a random equity against the S&P 500
```
king-spy random
```
random with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy random -t=[time period]
```

### wsb ###
Returns the top tickers mentioned on the r/wallstreetbets subreddit (WSB) and the related sentiment for each
```
king-spy wsb
```

### sectors ###
Returns the performance of various sectors over a time period
```
king-spy sectors
```
sectors with the optional -t (Time) flag - Currently accepts "1M", "3M", "6M", "1Y", "3Y", and "YTD" (default)
```
king-spy sectors -t=[time period]
```
sectors with optional -t flag and the -s (Stock) flag
```
king-spy sectors -t=[time period] -s=[ticker]
```

### all ###
Runs the c2s, high, low, sectors, and news cmds for a single ticker
```
king-spy all [ticker]
```
all with the optional -c (Chart) flag, which will additionally launch a chart of your specified equity in a new instance of your default browser
```
king-spy all [ticker] -c
```

# Installing
1. Clone the repo
2. Navigate to Alpaca at alpaca.markets
3. Create an account and generate paper trading API tokens
4. Store the tokens somewhere safe
5. Navigate back to your cloned repo
6. Create a .env file with the following entries:
```
APCA_API_KEY_ID=YourAlpacaAPIKeyID
APCA_API_SECRET_KEY=YourAlpacaSecretAPIKeyID
ENDPOINT=https://paper-api.alpaca.markets
```
7. Save your file and navigate to the config/config.go file
8. Change the viper.AddConfigPath("placeholder") in the Init func to the absolute path of your Alpaca .env file
9. Save and run the following for your cloned repo:
```
go build && go install
```
10. The application should now be added to your GOPATH and you should be able to run it from any directory on your system

# Recommendation
For ease of use, open your .bashrc or .zshrc file with your favorite text editor and add the following alias:

```
alias ks='king-spy'
```

# Disclaimer
The information provided through this application is not intended as, and shall not be understood or construed as, financial advice.

I am not an attorney, accountant or financial advisor, nor am I holding myself out to be, and the information provided by this application is not a substitute for financial advice from a professional who is aware of the facts and circumstances of your individual situation.
