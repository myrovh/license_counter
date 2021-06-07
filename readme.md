# Calculate Licenses Function

Given an input function will calculate the required license number for a given application (default 374). A single license is valid for one desktop and one laptop. So total required licenses can be calculated simply by counting each unique computer and summing the total desktops and laptops for each user, whichever number is higher is the total required licenses.

For ease of demonstration this repo has a 'main' function that acts as a command line interface to 'CalculateTotalLicenses'. CalculateTotalLicenses is to be integrated in a larger application and accepts an io.Reader and a applicationID which represents the application it should be counting from the csv. Any lines that don't match the applicationID are skipped when parsing the csv file.

## NOTE

The tests we're written to make use of all the provided sample data including the `sample-large.csv` however to save download time this file isn't included in the repo. To successfully run the benchmark without it skipping you can add the large sample file to the `example_files` folder.