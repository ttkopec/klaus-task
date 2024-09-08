# klaus-task

> [!WARNING]
> Even though the solution is complete and the basic functionality is delivered according to the task description, the solution for sure can be further improved and polished. Please see the [TODO](#todo) section for more details.

## Overview

The service exposes gRPC API (by default at port 5001).

Schema of the API can be viewed [here](./pkg/protos/score_service.proto).

## How to run?

```bash
make run
```
or using `docker`:
```bash
make docker_build
make docker_run
```

## Run tests

```
make test
```

## Example API calls
1. method: `ScoreService.GetAggregatedCategoryScores`
    - payload#1 - period longer than 1 month:
    ```json
    {
        "period": {
            "end": {
                "seconds": "1565992800"
            },
            "start": {
                "seconds": "1563228000"
            }
        }
    }
    ```
    - payload#2 - period shorter than 1 month:
    ```json
    {
        "period": {
            "end": {
                "seconds": "1563314400"
            },
            "start": {
                "seconds": "1563228000"
            }
        }
    }
    ```
1. method: `ScoreService.GetScoresByTicket`
    - payload: 
    ```json
    {
        "period": {
            "end": {
                "seconds": "1563228000"
            },
            "start": {
                "seconds": "1563141600"
            }
        }
    }
    ```

1. method: `ScoreService.GetOverallQualityScore`
    - payload: 
    ```json
    {
        "period": {
            "end": {
                "seconds": "1563228000"
            },
            "start": {
                "seconds": "1563141600"
            }
        }
    }
    ```

1. method: `ScoreService.GetPeriodOverPeriodScoreChange`
    - payload:
    ```json
    {
        "currentPeriod": {
            "end": {
                "seconds": "1563314400"
            },
            "start": {
                "seconds": "1563228000"
            }
        },
        "previousPeriod": {
            "end": {
                "seconds": "1563228000"
            },
            "start": {
                "seconds": "1563141600"
            }
        }
    }
    ```

## TODO

1. Dockerfile can be optimized by for example:
    - using multi stage build and introducing builder + runner layers so that the resulting image is not that heavy
1. add more tests
    - due to lack of time, I only created a few simple tests for [server_test.go](./internal/api/server_test.go), but for sure better test coverage is needed for the rest of the codebase
1. more comments and documentation