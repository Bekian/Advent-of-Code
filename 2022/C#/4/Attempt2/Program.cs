namespace Attempt2;
// this second attempt converts the input into a string of numbers that represents the range
// then compares range values to see if they overlap
// this attempt targets the second part of the problem

class Program
{
    static void Main(string[] args)
    {
        string[] content = [];
        try
        {
            content = File.ReadAllLines("test2.txt");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
        }
        // create a list of lists of strings that represent both ranges in each input row
        List<List<string>> ranges = [[]];
        // TODO: change this to an array of ints or strings?

        // assign the positions
        for (int i = 0; i < content.Length; i++)
        {
            string[] parts = content[i].Split(',');
            string[] firstParts = parts[0].Split('-');
            string[] secondParts = parts[1].Split('-');

            ranges.Add([firstParts[0], firstParts[1], secondParts[0], secondParts[1]]);
        }
        // remove the skill issue
        ranges.RemoveAt(0);

        // process the input from rows of numeric strings to sequential arrays of strings
        // ex: 1 - 3 would become "1", "2", "3"
        foreach (var row in ranges)
        {
            // parse the first range
            Console.WriteLine(string.Join(", ", row));
            int lowerBound = int.Parse(row[0]);
            int upperBound = int.Parse(row[1]);
            // create a string range with each digit contained as a string
            List<string> range = [];
            for (int i = lowerBound; i <= upperBound; i++)
            {
                range.Add($"{i}");
            }
            Console.WriteLine(string.Join(", ", range));
            // parse the second range
            int lowerBound2 = int.Parse(row[2].ToString());
            int upperBound2 = int.Parse(row[3].ToString());
            List<string> range2 = [];
            for (int i = lowerBound2; i <= upperBound2; i++)
            {
                range2.Add($"{i}");
            }
            Console.WriteLine(string.Join(", ", range2));
        }
        // need to compare range values to each other
    }
}
