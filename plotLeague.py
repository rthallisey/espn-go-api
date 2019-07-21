#!/usr/bin/env python

import sys
import os
import json

PROJECT_ROOT = os.path.abspath(os.path.join(
        os.path.dirname(os.path.realpath(__file__))))
if PROJECT_ROOT not in sys.path:
    sys.path.insert(0, PROJECT_ROOT)

#import plot

import plotly
import plotly.plotly as plot
import plotly.graph_objs as go


RESULTS_DIR="results"

with open(RESULTS_DIR + "/BenchPoints.json") as json_file:
    data = json.load(json_file)

    benchGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Weekly Bench Points Scored")]

    plotly.offline.plot(benchGraph, filename='weekly-bench-points.html')

sys.exit(0)

with open(RESULTS_DIR + "/BenchPoints.json") as json_file:
    data = json.load(json_file)
plot_data['filename'] = "avg-points-position.html"
plot_data['title'] = "Average Weekly Points Per Position from an Active Roster"
Average_points_plot = plot.Plot(plot_data)
#Average_points_plot.plot_offline()
#Average_points_plot.plot()

plot_data['filename'] = "team_mvp_value.html"
plot_data['title'] = "Team MVP value"
MVP_win_changes = plot.Plot(plot_data)
MVP_win_changes.bar([str(v) for v in mvp_value.keys()],
                        [v[0] for v in mvp_value.values()],
                        [v[1] for v in mvp_value.values()])
#MVP_win_changes.plot_offline()
# MVP_win_changes.plot()

plot_data['filename'] = "season-bench-points.html"
plot_data['title'] = "Total Bench Points"
Season_bench_points = plot.Plot(plot_data)
Season_bench_points.bar(pts.keys(), pts.values())
#Season_bench_points.plot_offline()
# Season_bench_points.plot()

plot_data['filename'] = "most-points-from-players.html"
plot_data['title'] = "Most Points from a Player from an Active Roster"
Season_bench_points = plot.Plot(plot_data)
Season_bench_points.bar(team_names, best_player_score, best_player)
#Season_bench_points.plot_offline()
# Season_bench_points.plot()
