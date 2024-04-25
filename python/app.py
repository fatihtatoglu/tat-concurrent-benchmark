import asyncio
import sys


async def perform_task():
    await asyncio.sleep(10)


async def main(task_count):
    tasks = []

    for taskId in range(task_count):
        task = asyncio.create_task(perform_task())
        tasks.append(task)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    if len(sys.argv) > 1:
        task_count = int(sys.argv[1])
    else:
        task_count = 10000

    asyncio.run(main(task_count))
    print("All tasks are finished.")
