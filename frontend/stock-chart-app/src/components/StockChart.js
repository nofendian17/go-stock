import React, { useEffect, useState, useRef } from 'react';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  TimeScale
} from 'chart.js';
import 'chartjs-adapter-date-fns';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  TimeScale
);

const StockChart = ({ ticker }) => {
  const [chartData, setChartData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      // Ganti dengan tanggal yang dinamis jika perlu
      const startDate = '2023-01-01';
      const endDate = '2023-12-31';

      try {
        const response = await fetch(`/api/v1/stock/historical?stock_code=${ticker}&start_date=${startDate}&end_date=${endDate}`);
        const data = await response.json();

        if (data) {
          setChartData({
            labels: data.map(item => new Date(item.date)),
            datasets: [
              {
                label: `${ticker} Close Price`,
                data: data.map(item => item.close),
                borderColor: 'rgb(75, 192, 192)',
                tension: 0.1,
              },
            ],
          });
        }
      } catch (error) {
        console.error("Error fetching stock data:", error);
      }
    };

    if (ticker) {
      fetchData();
    }
  }, [ticker]);

  if (!chartData) {
    return <div>Loading chart...</div>;
  }

  return (
    <div>
      <h2>{ticker} Stock Price</h2>
      <Line
        data={chartData}
        options={{
          scales: {
            x: {
              type: 'time',
              time: {
                unit: 'month'
              }
            }
          }
        }}
      />
    </div>
  );
};

export default StockChart;
