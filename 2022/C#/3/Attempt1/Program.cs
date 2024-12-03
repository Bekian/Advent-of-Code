using System.Diagnostics;
using System.Diagnostics.Metrics;

namespace Attempt1;

class Program
{
    static void Main(string[] args)
    {
        string[] content = [];
        try
        {
            content = File.ReadAllLines("input.txt");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
        }
        // use a hashmap here to define keys (letters) and values (ints)

        // get the string that will be used to compute the sum of all dupelicate char values
        string sumString = "";
        // check each line
        foreach (var line in content)
        {
            // find midpoint
            int midpoint = line.Length / 2;
            // get first half and second half of each line
            string firstHalf = line[..midpoint];
            string secondHalf = line[midpoint..];
            // loop through each character in the first half
            foreach (char c in firstHalf)
            {
                // if that letter is in the second half
                if (secondHalf.Contains(c))
                {
                    // add it to the sumString
                    sumString += c;
                    break;
                }
            }
        }
        // Console.WriteLine(sumString);

        // we need an alphabet for indexing
        string alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
        // Console.WriteLine(alphabet.Length);

        // compute the values of the sumString
        int sum = 0;
        foreach (char c in sumString)
        {
            // find the index of the character in the double alphabet
            int index = alphabet.IndexOf(c);
            // add 1 and the index to get the correct value
            sum += index + 1;
        }
        // first part answer:
        Console.WriteLine(sum);


        // for the second part we need to make every 3 lines a group of elf bags
        // we need to first find the common letter in every 3 lines
        // first group the first 3 rows 
        List<List<string>> groups = [[]];
        List<string> tempGroup = [];
        int counter = 1;
        foreach (var line in content)
        {
            tempGroup.Add(line);
            if (counter == 3)
            {
                groups.Add(tempGroup);
                counter = 0;
                tempGroup = [];
            }
            counter++;
        }
        groups.RemoveAt(0);

        // then find the common letter in each group and add it to the sumstring
        sumString = "";
        foreach (var group in groups)
        {
            foreach (char c in group[0])
            {
                if (group[1].Contains(c) && group[2].Contains(c))
                {
                    sumString += c;
                    break;
                }
            }
        }

        // compute the values of the sumString again
        sum = 0;
        foreach (char c in sumString)
        {
            int index = alphabet.IndexOf(c);
            sum += index + 1;
        }
        // second part answer:
        Console.WriteLine(sum);


    }
}
