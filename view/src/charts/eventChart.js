import { Bar } from 'vue-chartjs';

export default {
  extends: Bar,
  props: ['sourceData', 'options'],
  mounted() {
    let labels = [];
    let chartData = [];
    for(let data of this.sourceData) {
      labels.push(data.date);
      chartData.push(data.eventCount);
    }
    this.renderChart({
      labels: labels,
      datasets: [
        {
          label: 'Event',
          backgroundColor: '#0F733C',
          data: chartData,
        },
      ]
    },  {responsive: true, maintainAspectRatio: false});
  }
}