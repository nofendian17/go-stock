import React, { useState } from 'react';
import StockChart from './components/StockChart';
import './App.css';

function App() {
  const [ticker, setTicker] = useState("BBCA");
  const [inputValue, setInputValue] = useState("BBCA");

  const handleFetch = () => {
    setTicker(inputValue.toUpperCase());
  };

  return (
    <div className="App">
      <h1>Stock Market Chart</h1>
      <div>
        <input
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Enter stock ticker (e.g., BBCA)"
        />
        <button onClick={handleFetch}>Fetch Data</button>
      </div>
      <StockChart ticker={ticker} />
    </div>
  );
}

export default App;
