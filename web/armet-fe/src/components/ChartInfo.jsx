import React from 'react';

const ChartInfo = (props) => {
    return (
        <div>
        <h2>{props.chart.chart.metadata.name}</h2>
        <p>Chart version: {props.chart.chart.metadata.version}</p>
        <p>Application version: {props.chart.chart.metadata.appVersion}</p>
        <p>Status: {props.chart.info.status}</p>

        </div>          
    );
};

export default ChartInfo;