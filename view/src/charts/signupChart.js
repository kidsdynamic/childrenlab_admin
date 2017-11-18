import { Bar } from 'vue-chartjs';

export default {
  extends: Bar,
  props: ['sourceData', 'options'],
  mounted() {
    let labels = [];
    let chartData = [];
    for(let data of this.sourceData) {
      labels.push(data.Date);
      chartData.push(data.SignupCount);
    }
    this.renderChart({
      labels: labels,
      datasets: [
        {
          label: 'Signup',
          backgroundColor: '#FD733D',
          data: chartData,
        },
      ]
    },  {responsive: true, maintainAspectRatio: false});
  }
}