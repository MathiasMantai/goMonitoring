var data = {
    labels: ['free', 'in use'],
    datasets: [{
        data: [100,0], // Angenommene Datenwerte
        backgroundColor: ['#38B3F8', '#ddd']
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
    },
    elements: {
        center: {
            text: 'Hallo Welt'
        }
    }
};


var ctx = document.getElementById('cpuUsage').getContext('2d')
let cpuPercent = document.getElementById('cpuUsage-percent')
let text = "t"
let donutChart = new Chart(ctx, {
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
    fetch("/cpuUsage")
    .then(response => response.json())
    .then(
        (data) => {
            console.log(data)
            cpu = data[0].toFixed(2)
            updateChart(donutChart, [parseFloat(100-data), data])
            cpuPercent.innerHTML = data + "%"
        }
    )
}