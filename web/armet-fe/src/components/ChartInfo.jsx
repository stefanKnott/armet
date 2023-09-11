import React from 'react';

const ChartInfo = (props) => {
    return (
        <div>
        <div >
        <h1 className="chartName">{props.chart.chart.metadata.name}</h1>
        {props.chart.info.status === "deployed" &&
            <div className="deployed"/>
        }
         {props.chart.info.status != "deployed" &&
            <div className="failed"/>
        }
        </div>
        <p>Chart version: {props.chart.chart.metadata.version}</p>
        <p>Application version: {props.chart.chart.metadata.appVersion}</p>
        <p>Status: {props.chart.info.status}</p>

        </div>          
    );
};

export default ChartInfo;