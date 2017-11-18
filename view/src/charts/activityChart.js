import { Bar } from 'vue-chartjs';

export default {
  extends: Bar,
  props: ['sourceData', 'options'],
  mounted() {
    let labels = [];
    let activityCount = [];
    let indoor = [];
    let outdoor = [];
    let userCount = [];
    for(let data of this.sourceData) {
      labels.push(data.Date);
      activityCount.push(data.ActivityCount);
      indoor.push(data.IndoorSteps);
      outdoor.push(data.OutdoorSteps);
      userCount.push(data.UserCount);
    }
    this.renderChart({
      labels: labels,
      datasets: [
        {
          label: 'Activity Count',
          backgroundColor: '#FD733D',
          data: activityCount,
        },
        {
          label: 'Indoor Steps',
          backgroundColor: '#0074D9',
          data: indoor,
        },
        {
          label: 'Outdoor Steps',
          backgroundColor: '#7FDBFF',
          data: outdoor,
        },
        {
          label: 'User Count',
          backgroundColor: '#39CCCC',
          data: userCount,
        },
      ]
    },  {responsive: true, maintainAspectRatio: false});
  }
}