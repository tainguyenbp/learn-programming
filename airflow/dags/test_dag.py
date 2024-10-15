from datetime import datetime
from airflow import DAG
from airflow.decorators import task
from airflow.operators.bash import BashOperator

default_args = {
    'start_date': datetime(2023, 1, 1),  # start_date is required but won't affect as schedule_interval is None
    'retries': 0,                        # No retries in this case
}

with DAG(
    dag_id='adhoc_dag_main',
    default_args=default_args,
    catchup=False,                    # Disable backfilling or catching up
    schedule_interval=None,           # No schedule, ad-hoc execution
) as dag:
    # Tasks are represented as operators
    hello = BashOperator(task_id="hello", bash_command="echo hello8")

    @task()
    def airflow():
        print("airflow")

    # Set dependencies between tasks
    hello >> airflow()
