import { Sidebar, Menu, MenuItem, useProSidebar, SubMenu } from 'react-pro-sidebar';
import React, {useState, useEffect} from 'react';
import ChartInfo from './ChartInfo.jsx'


const Layout = () => {

    const [chartsByNamespace, setChartsByNamespace] = useState();
    const [isLoading, setIsLoading] = useState(true);
    const [displayChart, setDisplayChart] = useState(false)
    const [chartView, setChartView] = useState(null)
    const [chart, setChart] = useState(null);


    useEffect(() => {

    const fetchData = () => {
      fetch('http://localhost:8080/api/v1/helm/releases')
      .then(response => {
          return response.json()
      })
      .then(data => {
        setChartsByNamespace(data["releases"]);
        console.log(data);
        setIsLoading(false);
      })
    };

    fetchData();
  }, []);

  const getNumDeployed = (charts) => {
    let numDeployed = 0;
    charts.map((chart) => {
      if (chart.info.status == "deployed"){
        numDeployed++;
      }
    })

    return numDeployed;
  }

  const displayChartCard = (chart) => {
      setChart(chart)
      setDisplayChart(true)
  }

  const getNamespaceLabel = (namespace, charts) => {
    let namespaceLabel = namespace;
    namespaceLabel = namespaceLabel + " (" + getNumDeployed(charts) + "/" +charts.length + ")";
    return namespaceLabel;
  }

  const Namespaces = ({namespaces}) => (
    <div>
      {Object.entries(namespaces).map(([namespace, charts]) => (
      <SubMenu key={namespace} label={getNamespaceLabel(namespace,charts)}>
          {charts.map((chart) => (
          <MenuItem onClick={() => displayChartCard(chart)} active="true" key={chart.name}>{chart.name}</MenuItem>
          ))}
      </SubMenu>
      ))}
    </div>  
  ) 

    return (
      <div style={{ display: 'flex', height: '100%' }}>
        <Sidebar>
        {isLoading ? (
          <p>loading</p>
        ) : (
          <Menu>
           {Object.entries(chartsByNamespace).map(([cluster, namespaces]) => (
             <SubMenu key={cluster} label={cluster}>
              <Namespaces namespaces={namespaces}/>
            </SubMenu>
            ))} 
          </Menu>
          )}
        </Sidebar>

        {displayChart ? (
          <ChartInfo chart={chart}/>
          ) : (
            <p></p>
        )}
      </div>              
    );
};

export default Layout;