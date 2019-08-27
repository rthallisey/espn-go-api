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

def buildTrace(title, X, Y):
    fig = go.Figure()
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

    return fig

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


def getWins(teamOwner):
    if teamOwner == "Carlo Masakayan":
        return 3 * 100
    elif teamOwner == "Dan Rodenbach":
        return 6 * 100
    elif teamOwner == "Daniel McCarthy":
        return 7 * 100
    elif teamOwner == "David Hastie":
        return 12 * 100
    elif teamOwner == "Joe Lepera":
        return 5 * 100
    elif teamOwner == "John Gentile":
        return 8 * 100
    elif teamOwner == "Kyle Johnson":
        return 7 * 100
    elif teamOwner == "Matt Dunn":
        return 6 * 100
    elif teamOwner == "Michael Fitzgerald":
        return 10 * 100
    elif teamOwner == "Michael Hofmann":
        return 7 * 100
    elif teamOwner == "Ryan Hallisey":
        return 8 * 100
    elif teamOwner == "Chris BRos":
        return 5 * 100

with open(RESULTS_DIR + "/BenchPoints.json") as json_file:
    data = json.load(json_file)

    benchGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Total Bench Points Scored")]

    buildPlot("Total Bench Points Scored", "total-bench-points", benchGraph, "Team Owner", "Points")

    benchVSWinsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Total Bench Points Scored"),
        go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(getWins(v)) for v in data.keys()],
        name="Wins*100")]

    buildPlot("Bench Points and Wins", "total-bench-points-vs-wins", benchVSWinsGraph, "Team Owner", "Points/Wins x 100")

with open(RESULTS_DIR + "/MVPCount.json") as json_file:
    data = json.load(json_file)

    MVPCountGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Unique MVPs")]

    buildPlot("Unique MVPs", "unique-mvps", MVPCountGraph, "Team Owner", "# of MVPs")

with open(RESULTS_DIR + "/NormalizedMVPsPerWin.json") as json_file:
    data = json.load(json_file)

    NormMVPsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="How Hard Was It to Win")]

    buildPlot("How Hard Was It to Win", "normalized-mvps", NormMVPsGraph, "Team Owner", "Difficulty Score")

with open(RESULTS_DIR + "/playerMostPoints.json") as json_file:
    data = json.load(json_file)

    MostPointsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="MVP by Points")]

    buildPlot("MVP by Points", "team-mvp-points", MostPointsGraph, "Team Owner", "Points")

with open(RESULTS_DIR + "/AvgMVPs.json") as json_file:
    data = json.load(json_file)

    MostPointsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Average MVPs")]

    buildPlot("Average MVP", "avg-mvps", MostPointsGraph, "Team Owner", "MVPs")

t = buildTrace("Team Points Per Position", "Team Owner", "Points")
for filename in os.listdir(RESULTS_DIR + "/TeamPtsPerPosition/"):
    with open(RESULTS_DIR + "/TeamPtsPerPosition/" + filename) as json_file:
        data = json.load(json_file)
        t.add_trace(go.Scatter(x=[str(v) for v in data.keys()], y=[str(v) for v in data.values()],
                               mode='lines+markers',
                               name=filename.replace(".json", "")))

d = plotly.offline.plot(t, filename="team-pts-per-position.html", output_type='div')
with open(os.path.join(BLOG_PATH, "static/html", "team-pts-per-position.html"), "w") as f:
    f.write(d)


MVPsByWinsGraph = []
for filename in os.listdir(RESULTS_DIR + "/TeamMVPByWin/"):
    names = ""
    wins = ""
    with open(RESULTS_DIR + "/TeamMVPByWin/" + filename) as json_file:
        data = json.load(json_file)

        for k in data.keys():
            if names is "" :
                names += k
            else:
                names += " & " + k

        for v in data.values():
            wins = v
            break

    MVPsByWinsGraph.append(go.Bar(
        x=[names],
        y=[wins],
        name=filename.replace(".json", "")))

buildPlot("MVPs by Wins", "mvps-by-wins", MVPsByWinsGraph, "Team Owner", "Games Won")

sys.exit(0)
