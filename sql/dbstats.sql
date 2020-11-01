select
  event_time, stats_time, stats_date, stats_date_next, date(event_time) > stats_date as should_update, -- select date_part('hour', current_time)  > 1
  now() > (stats_date + INTERVAL '10minute') and stats_date_next < now() as should_update_normal
from (
  select
    event_time, stats_time,  
    cast(date(stats_time) as timestamptz) + INTERVAL '1day'  as stats_date, 
    cast(date(stats_time) + INTERVAL '2day' as timestamptz) as stats_date_next
  from 
  (
    select event_time, coalesce(stats_time_current, event_time_init - INTERVAL '1day') as stats_time
    from
      (
        select
          max(time) as event_time,
          min(time) as event_time_init
        from
          contract_events
        where
          netid = 3
          and contract = 'token_staking'
      ) a
      cross join (
        select
          max(time) as stats_time_current
        from
          operatorstats
        where
          netid = 3
          and contract = 'token_staking'
      ) b
  ) c
)d