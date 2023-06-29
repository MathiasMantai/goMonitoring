if(window.location.href.includes("/network"))
{
    const latencies = [];

    // Erstelle ein neues Chart.js-Diagramm
    const ctx = document.getElementById('latencyChart').getContext('2d');
    const chart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: ['1', '2', '3', '4', '5', '6', '7', '8', '9', 
                 '10', '11','12','13','14','15','16','17','18','19','20'
                ],
        datasets: [
        {
            label: 'Latenzzeit (ms)',
            data: latencies,
            backgroundColor: '#38B3F8',
            borderColor: 'rgba(0, 123, 255, 1)',
            borderWidth: 1,
            fill: true
        }
        ]
    },
    options: {
        animation: {
            duration: 0,
        },
        hover: {
            animationDuration: 0,
        },
        responsiveAnimationDuration: 0,
        responsive: true,
        scales: {
            y: {
                beginAtZero: true,
                grid: {
                    display: false
                }
            },
            x: {
                offset: true,
                reverse: true,
                ticks: {
                    display: false
                },
                grid: {
                    display: false
                }
            }
        },
        plugins: {
            legend: {
                display: false
            }
        }
    }
    });

    // Render das Diagramm
    chart.render()

    setInterval(() => {
        updateLatencyChart(chart)
    }, 2000)

    function updateLatencyChart(chart)
    {
        let data = chart.data.datasets[0].data
        if(data.length >= 20)
            data.pop()

        data.unshift(parseInt(Math.random() * 100))

        chart.update()
    }
}
else if(window.location.href.includes("/container"))
{

}
else
{
    var data = {
        labels: ['free', 'in use'],
        datasets: [{
            data: [100,0], // Angenommene Datenwerte
            backgroundColor: ['#ddd', '#38B3F8']
        }]
    };

    // Konfiguration des Charts
    var options = {
        cutout: '80%', // Der Prozentsatz des inneren Lochs (0-100)
        responsive: false,
        maintainAspectRatio: false,
        plugins: {
            customCanvasBackgroundColor: {
                color: '#38B3F8'
            },
            legend: {
                display: false
            },
        }
    };


    var ctx = document.getElementById('cpuUsage').getContext('2d')
    var ctx2 = document.getElementById('memoryUsage').getContext('2d')
    let cpuPercent = document.getElementById('cpuUsage-percent')
    let memoryPercent = document.getElementById('memory-percent')
    let donutChart = new Chart(ctx, {
        type: 'doughnut',
        data,
        options
    })


    let donutChart2 = new Chart(ctx2, {
        type: 'doughnut',
        data,
        options
    })

    getCpuData()

    setInterval(() => {
        getCpuData()
    },2000)


    function updateChart(chart, newData)
    {
        chart.data.datasets[0].data = [];
        newData.forEach(dataItem => {
            chart.data.datasets[0].data.push(dataItem)
        })
        chart.update()
    }

    function getCpuData()
    {
        fetch("/cpu")
        .then(response => response.json())
        .then(
            (data) => {
                console.log(data)
                cpu = data.CpuUsage.toFixed(2)
                vMemory = data.VirtualMemory.toFixed(2)
                updateChart(donutChart, [parseFloat(100-cpu), cpu])
                updateChart(donutChart2, [parseFloat(100-vMemory), vMemory])
                cpuPercent.innerHTML = cpu + "%"
                memoryPercent.innerHTML = vMemory + "%"
            }
        )
    }
}