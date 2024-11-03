<template>
    <div>
      <h2>网站总访问量: {{ totalVisits }}</h2>
      <line-chart :data="chartData"></line-chart>
    </div>
  </template>
  
  <script>
  import LineChart from '@/components/LineChart';
  
  export default {
    name: "AdminPage",
    components: {
      LineChart
    },
    data() {
      return {
        totalVisits: 0,
        chartData: {
          labels: [],
          series: []
        }
      };
    },
    mounted() {
      this.generateData();
    },
    methods: {
      generateData() {
        const labels = [];
        const series = [];
        let total = 0;
        const today = new Date();
        for (let i = 29; i >= 0; i--) {
          const day = new Date(today.getFullYear(), today.getMonth(), today.getDate() - i);
          const visits = Math.floor(Math.random() * 100) + 10;
          labels.push(`${day.getDate()}/${day.getMonth() + 1}`);
          series.push(visits);
          total += visits;
        }
  
        this.totalVisits = total;
        this.chartData = { labels, series };
      }
    }
  }
  </script>
  