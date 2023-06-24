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