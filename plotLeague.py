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
import chart_studio.plotly as plot
import plotly.graph_objects as go


RESULTS_DIR=PROJECT_ROOT + "/results"
GO_PATH=os.environ['GOPATH']
BLOG_PATH=os.path.join(GO_PATH, "src/github.com/rthallisey/blog")

def buildPlot(title, filename, data, X, Y):
    fig = go.Figure(data=data)
    fig.update_layout(
        title=go.layout.Title(
            text=title,
            font=dict(
                size=32
            ),
            xref="paper",
            x=0.5
        ),
        xaxis=go.layout.XAxis(
            title=go.layout.xaxis.Title(
                text=X,
                font=dict(
                    family="Courier New, monospace",
                    size=20,
                    color="#7f7f7f",
                )
            )
        ),
        yaxis=go.layout.YAxis(
            title=go.layout.yaxis.Title(
                text=Y,
                font=dict(
                    family="Courier New, monospace",
                    size=20,
                    color="#7f7f7f",
                )
            )
        )
    )
    d = plotly.offline.plot(fig, filename=filename, output_type='div')
    with open(os.path.join(BLOG_PATH, "static/html", filename+".html"), "w") as f:
        f.write(d)

with open(RESULTS_DIR + "/BenchPoints.json") as json_file:
    data = json.load(json_file)

    benchGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Weekly Bench Points Scored")]

    buildPlot("Weekly Bench Points Scored", "weekly-bench-points", benchGraph, "Team Owner", "Points")

with open(RESULTS_DIR + "/MVPCount.json") as json_file:
    data = json.load(json_file)

    AvgMVPsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Average MVPs")]

    buildPlot("Average Number of MVPs", "avg-mvps", AvgMVPsGraph, "Team Owner", "# of MVPs")

sys.exit(0)
