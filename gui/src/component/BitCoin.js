import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Line } from 'react-chartjs-2'

const BitCoin = () => {
    const [chartData, setChartData] = useState({});
    const chartOptions = {
        scales: {
            xAxes: [{
                display: true,
                type: 'time',
                time: {
                    unit: 'week'
                }
            }],
            yAxes: [{
                ticks: {
                    beginAtZero: true,
                    callback: function (value, index, values) {
                        if (parseInt(value) >= 1000) {
                            return '$' + value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
                        } else {
                            return '$' + value;
                        }
                    }
                }
            }]
        }
    }
    const chart = () => {
        let date = []
        let rate = []
        axios.get('https://api.esquire.jessup.info/v1/answer/1/BTC/2018')
            .then(res => {
                for (const dataObj of res.data) {
                    date.push(Date.parse(dataObj.timestamp));
                    rate.push(dataObj.rate);
                }
                setChartData({
                    labels: date,
                    datasets: [
                        {label: "rate", data: rate},
                    ]
                })
            })
            .catch(err => {
                console.log(err)
            });
    }

    useEffect(() => {
        chart()
    }, [])

    return (
        <div className="App">
            <h1>BitCoin</h1>
            <h2>Year to Date, by Week | 2018</h2>
            <div><Line data={chartData} options={chartOptions} /></div>
        </div>
    )
}

export default BitCoin;