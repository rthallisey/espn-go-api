#!/usr/bin/env python

import sys
import os
import json

PROJECT_ROOT = os.path.abspath(os.path.join(
        os.path.dirname(os.path.realpath(__file__))))
if PROJECT_ROOT not in sys.path:
    sys.path.insert(0, PROJECT_ROOT)

import plotly
import chart_studio.plotly as plot
import plotly.graph_objects as go

import utils
import wins

RESULTS_DIR=PROJECT_ROOT + "/../results"
GO_PATH=os.environ['GOPATH']
BLOG_PATH=os.path.join(GO_PATH, "src/github.com/rthallisey/blog")
STATIC_HTML_PATH=os.path.join(BLOG_PATH, "static/html")


def totalBenchPointsScored(filename):
    data = utils.ReadResult(RESULTS_DIR + "/BenchPoints.json")
    benchGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Total Bench Points Scored")
    ]
    g = utils.BuildPlot("Total Bench Points Scored", filename, benchGraph, "Team Owner", "Points")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def benchVsWinsGraph(filename):
    data = utils.ReadResult(RESULTS_DIR + "/BenchPoints.json")
    benchVSWinsGraph = [
        go.Bar(
            x=[str(v) for v in data.keys()],
            y=[str(v) for v in data.values()],
            name="Total Bench Points Scored"),
        go.Bar(
            x=[str(v) for v in data.keys()],
            y=[str(wins.getWins(v)) for v in data.keys()],
            name="Wins*100")
    ]
    g = utils.BuildPlot("Bench Points and Wins", filename, benchVSWinsGraph, "Team Owner", "Points/Wins x 100")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def mvpCount(filename):
    data = utils.ReadResult(RESULTS_DIR + "/MVPCount.json")
    MVPCountGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Unique MVPs")
    ]
    g = utils.BuildPlot("Unique MVPs", filename, MVPCountGraph, "Team Owner", "# of MVPs")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def normalizedMVPs(filename):
    data = utils.ReadResult(RESULTS_DIR +  "/NormalizedMVPsPerWin.json")

    NormMVPsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="How Hard Was It to Win")]

    g = utils.BuildPlot("How Hard Was It to Win", filename, NormMVPsGraph, "Team Owner", "Difficulty Score")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def teamMvpPoints(filename):
    data = utils.ReadResult(RESULTS_DIR +  "/playerMostPoints.json")

    MostPointsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="MVP by Points")
    ]
    g = utils.BuildPlot("MVP by Points", filename, MostPointsGraph, "Team Owner", "Points")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def avgMvps(filename):
    data = utils.ReadResult(RESULTS_DIR +  "/AvgMVPs.json")

    MostPointsGraph = [go.Bar(
        x=[str(v) for v in data.keys()],
        y=[str(v) for v in data.values()],
        name="Average MVPs")
    ]
    g = utils.BuildPlot("Average MVP", filename, MostPointsGraph, "Team Owner", "MVPs")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def teamPointsPerPosition(filename):
    t = utils.BuildTrace("Team Points Per Position", "Team Owner", "Points")
    for filename in os.listdir(RESULTS_DIR + "/TeamPtsPerPosition/"):
        data = utils.ReadResult(RESULTS_DIR + "/TeamPtsPerPosition/" + filename)
        t.add_trace(go.Scatter(x=[str(v) for v in data.keys()], y=[str(v) for v in data.values()],
                               mode='lines+markers',
                               name=filename.replace(".json", ""))
                    )
    g = plotly.offline.plot(t, filename="team-pts-per-position.html", output_type='div')
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def mvpsByWins(filename):
    MVPsByWinsGraph = []
    for filename in os.listdir(RESULTS_DIR + "/TeamMVPByWin/"):
        names = ""
        wins = ""
        data = utils.ReadResult(RESULTS_DIR + "/TeamMVPByWin/" + filename)

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

    g = utils.BuildPlot("MVPs by Wins", "mvps-by-wins", MVPsByWinsGraph, "Team Owner", "Games Won")
    utils.PublishGraphToBlog(STATIC_HTML_PATH, filename, g)


def main():
    totalBenchPointsScored("total-bench-points")
    benchVsWinsGraph("total-bench-points-vs-wins")
    mvpCount("mvp-count")
    normalizedMVPs("normalized-mvps")
    teamMvpPoints("team-mvp-points")
    avgMvps("avg-mvps")
    teamPointsPerPosition("team-pts-per-position")
    mvpsByWins("mvps-by-wins")


if __name__ == "__main__":
    main()
