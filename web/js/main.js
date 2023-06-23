var data = {
    labels: ['Rot', 'Grün'],
    datasets: [{
        data: [30, 50], // Angenommene Datenwerte
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
    }
};

// Chart erstellen
var ctx = document.getElementById('chart').getContext('2d');
var donutChart = new Chart(ctx, {
    type: 'doughnut',
    data,
    options
});

// Chart erstellen
var ctx2 = document.getElementById('chart2').getContext('2d');
var donutChart = new Chart(ctx2, {
    type: 'doughnut',
    data,
    options
});

