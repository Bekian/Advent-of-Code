namespace Attempt1;

using System;
using System.IO;
using System.Linq;

class Program
{
    static void Main(string[] args)
    {
        // the main idea for the first part of the problem is to parse every group of numbers
        // then find the first group with the highest sum
        // the main idea for the 2nd part is to find the sum of the top 3 sums
        // the simplest way i thought of to do this was to get all the sums 
        // sort them in descending order and then get the first 3 and add them
        int highestSum = 0;
        string content = "";
        try
        {
            content = File.ReadAllText("input.txt");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
        }

        var lines = content.Split(new[] { Environment.NewLine }, StringSplitOptions.None);
        var currSum = 0;
        var nElfs = 0;
        List<int> allSums = [];
        foreach (var line in lines)
        {
            var attemptParse = int.TryParse(line, out var currNum);
            // if the parse failed, we know there was a gap, and that means a new elf
            if (!attemptParse)
            {
                // check if currsum is bigger than highestsum
                if (currSum > highestSum)
                {
                    // if it is the new highest sum is the current sum
                    highestSum = currSum;
                }
                // push the current sum onto the list of allSums
                allSums.Add(currSum);
                // reset curr sum, increment nElfs
                currSum = 0;
                nElfs++;
            }

            currSum += currNum;
        }
        // these following 2 lines are equivalent
        // List<int> sortedSums = allSums.OrderByDescending(n => n).ToList();
        List<int> sortedSums = [.. allSums.OrderByDescending(n => n)];
        int top3Sum = sortedSums[0] + sortedSums[1] + sortedSums[2];
        Console.WriteLine($"Highest Sum: {highestSum}, nelfs: {nElfs}, top3 sum: {top3Sum}");
    }
}
