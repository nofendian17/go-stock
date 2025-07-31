import React from 'react';
import Chart from 'react-apexcharts';
import type { StockSummaryResponse } from '../types/stock';
import type { ApexOptions } from 'apexcharts';

interface StockChartProps {
    summaries: StockSummaryResponse[];
    title?: string;
}

export const StockChart: React.FC<StockChartProps> = ({ summaries, title = 'Stock Price Chart' }) => {
    const summariesArray = Array.isArray(summaries) ? summaries : [];

    if (summariesArray.length === 0) {
        return <div className="text-center py-8 text-gray-500">No data available</div>;
    }

    const sortedSummaries = [...summariesArray].sort((a, b) =>
        new Date(a.date).getTime() - new Date(b.date).getTime()
    );

    const calculateSMA = (data: any[], period: number) => {
        const sma = [];
        for (let i = period - 1; i < data.length; i++) {
            const sum = data.slice(i - period + 1, i + 1).reduce((acc, val) => acc + val.close, 0);
            sma.push({
                x: new Date(data[i].date),
                y: sum / period
            });
        }
        return sma;
    };

    const sma20 = calculateSMA(sortedSummaries, 20);

    const seriesCandlestick = [
        {
            name: 'Candlestick',
            type: 'candlestick',
            data: sortedSummaries.map(summary => ({
                x: new Date(summary.date),
                y: [summary.open_price, summary.high, summary.low, summary.close]
            }))
        },
        {
            name: 'SMA 20',
            type: 'line',
            data: sma20,
            color: '#FF0000'
        }
    ];

    const seriesBar = [{
        name: 'Volume',
        data: sortedSummaries.map(summary => ({
            x: new Date(summary.date),
            y: summary.volume
        }))
    }];

    const optionsCandlestick: ApexOptions = {
        chart: {
            type: 'candlestick',
            height: 350
        },
        title: {
            text: title,
            align: 'left'
        },
        xaxis: {
            type: 'datetime'
        },
        yaxis: {
            tooltip: {
                enabled: true
            }
        }
    };

    const optionsBar: ApexOptions = {
        chart: {
            height: 160,
            type: 'bar',
            toolbar: {
                show: false
            },
            zoom: {
                enabled: false
            }
        },
        plotOptions: {
            bar: {
                columnWidth: '80%',
                colors: {
                    ranges: [{
                        from: -1000000000,
                        to: 0,
                        color: '#F15B46'
                    }, {
                        from: 1,
                        to: 1000000000,
                        color: '#FEB019'
                    }]
                }
            }
        },
        stroke: {
            width: 0
        },
        xaxis: {
            type: 'datetime',
            labels: {
                show: false
            }
        },
        yaxis: {
            labels: {
                show: false
            }
        }
    };

    return (
        <div className="bg-white rounded-lg shadow-md p-6">
            <Chart options={optionsCandlestick} series={seriesCandlestick} type="candlestick" height={350} />
            <Chart options={optionsBar} series={seriesBar} type="bar" height={160} />
        </div>
    );
};
