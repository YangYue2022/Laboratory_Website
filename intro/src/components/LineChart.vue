<template>
    <div ref="chartContainer" style="width: 100%; height: 400px;"></div>
  </template>
  
  <script>
  import * as echarts from 'echarts';
  
  export default {
    name: 'LineChart',
    props: ['data'],
    data() {
      return {
        chart: null,
      };
    },
    mounted() {
      this.chart = echarts.init(this.$refs.chartContainer);
      this.updateChart();
    },
    methods: {
      updateChart() {
        this.chart.setOption({
          title: {
            text: '每日访问量'
          },
          tooltip: {
            trigger: 'axis'
          },
          xAxis: {
            type: 'category',
            data: this.data.labels
          },
          yAxis: {
            type: 'value'
          },
          series: [{
            data: this.data.series,
            type: 'line',
            smooth: true
          }]
        });
      }
    },
    watch: {
      data: {
        deep: true,
        handler() {
          this.updateChart();
        }
      }
    }
  }
  </script>
  