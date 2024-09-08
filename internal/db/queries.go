package db

const aggregatedCategoryScoresQuery = `
WITH filtered_by_date AS (
  SELECT 
    rc.name AS category_name, 
    r.rating, 
    r.created_at 
  FROM 
    ratings r 
    JOIN rating_categories rc ON r.rating_category_id = rc.id 
  WHERE 
    r.created_at BETWEEN ? AND ?    
) 
SELECT 
  category_name, 
  CASE
    WHEN (JULIANDAY(?) - JULIANDAY(?)) <= 31 THEN strftime('%Y-%m-%d', created_at) -- Group by day if the date range is within a month
    ELSE strftime('%Y-%W', created_at) -- Group by week for periods longer than a month
  END AS period, 
  ROUND(
    SUM(
      cast(rating as real)
    ) / (
      count(*) * 5
    ) * 100, 
    2
  ) AS score 
FROM 
  filtered_by_date 
GROUP BY 
  category_name, period 
ORDER BY 
  period, category_name;
`

const scoresByTicketQuery = `
SELECT 
  r.ticket_id AS ticket_id, 
  ROUND(
    CAST(r.rating AS REAL) / 5 * 100
  ) AS score, 
  rc.name AS category 
FROM 
  ratings r 
  INNER JOIN rating_categories rc ON rc.id = r.rating_category_id 
WHERE 
  r.created_at BETWEEN ? AND ? 
ORDER BY 
  r.ticket_id, rc.name;
`

const overallQualityScoreQuery = `
SELECT 
  ROUND(
    SUM(
      CAST(r.rating AS REAL) / 5 * 100 * rc.weight
    ) / SUM(rc.weight), 
    2
  ) AS score 
FROM 
  ratings r 
  INNER JOIN rating_categories rc ON rc.id = r.rating_category_id 
WHERE 
  r.created_at BETWEEN ? AND ?;
`

const periodOverPeriodScoreChangeQuery = `
SELECT 
  COALESCE(
    ROUND(
      SUM(
        CAST(r.rating AS REAL) / 5 * 100 * rc.weight
      ) / SUM(rc.weight) - (
        SELECT 
          SUM(
            CAST(r1.rating AS REAL) / 5 * 100 * rc1.weight
          ) / SUM(rc1.weight) 
        FROM 
          ratings r1 
          INNER JOIN rating_categories rc1 ON rc1.id = r1.rating_category_id 
        WHERE 
          r1.created_at BETWEEN ? AND ?
      ), 
      2
    ), 
    0
  ) AS scoreChange 
FROM 
  ratings r 
  INNER JOIN rating_categories rc ON rc.id = r.rating_category_id 
WHERE 
  r.created_at BETWEEN ? AND ?;

`