Given an input (csv) application will calculate the required license number for a given application (default 374). A single license is valid for one desktop and one laptop. So total required licenses can be calculated simply by counting each unique computer and summing the total desktops and laptops whichever number is higher is the total required licenses.

Break program into stages. Parse csv into memory; use maps and structs to break csv down into user-> application -> {computer_id, computer_type}

pass go data into next function that will do the business calculation and figure out how many licenses each user needs.