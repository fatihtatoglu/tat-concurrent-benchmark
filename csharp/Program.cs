namespace Program
{
    using System;
    using System.Collections.Generic;
    using System.Threading.Tasks;

    public static class Program
    {
        public static async Task Main(string[] args)
        {
            int taskCount = 10000;
            if (args.Length > 0)
            {
                taskCount = int.Parse(args[0]);
            }

            List<Task> tasks = new List<Task>();
            for (int i = 0; i < taskCount; i++)
            {
                Task task = Task.Run(async () =>
                {
                    await Task.Delay(10 * 1000);
                });

                tasks.Add(task);
            }

            await Task.WhenAll(tasks);

            Console.WriteLine("All tasks are finished.");
        }
    }
}