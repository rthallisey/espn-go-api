import sys
import os
import json

import plotly
import plotly.graph_objects as go

def ReadResult(resultFile):
    with open(resultFile) as json_file:
        return json.load(json_file)


def BuildTrace(title, X, Y):
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


def BuildPlot(title, filename, data, X, Y):
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
    return d

def PublishGraphToBlog(blog, filename, graph):
    with open(os.path.join(blog, filename+".html"), "w") as f:
        f.write(graph)
