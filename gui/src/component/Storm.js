import React, {useState, useEffect} from 'react';
import axios from 'axios';
//import {Line} from 'react-chartjs-2'

const Storm = () => {
    const [chartData, setChartData] = useState({})

    const fetchStormData = () => {
        const stormDataAPI = 'https://api.esquire.jessup.info/v1/answer/3/Kentucky/2018'
        const getStormData = axios.get(stormDataAPI)
        axios.all([getStormData]).then(
            axios.spread((...allData) => {
                const stormData = allData[0]
                console.log(stormData)
            })
        )
    }

    useEffect(() => {
        fetchStormData()
    }, [])

    return(
        <div className="App">
            <h1>Storm Event Data</h1>
        </div>
    )
}

export default Storm;