import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Line } from 'react-chartjs-2'

const Storm = () => {
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
        let propertyDamage = []
        axios.get('https://api.esquire.jessup.info/v1/answer/3/Kentucky/2018')
            .then(res => {
                for (const dataObj of res.data) {
                    if (dataObj.Damage > 0) {
                        date.push(Date.parse(dataObj.Date));
                        propertyDamage.push(dataObj.Damage);
                    }
                }
                setChartData({
                    labels: date,
                    datasets: [
                        {label: "damage", data: propertyDamage}
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
            <h1>Storm Event Data</h1>
            <h2>Kentucky | 2018</h2>
            <div><Line data={chartData} options={chartOptions} /></div>
        </div>
    )
}

export default Storm;