# Ensure the 'ape' package is loaded for phylogenetic tree operations
if (!requireNamespace("ape", quietly = TRUE)) {
  install.packages("ape")
  library(ape)
} else {
  library(ape)
}

# Install the readxl package if it's not already installed
if (!require(readxl)) {
  install.packages("readxl")
  library(readxl)
} else {
  library(readxl)
}

# Under the "Session" menu, set the working directory to the source file location.

# First, we write a function completing a simple demo of how to draw a Newick tree
NewickDemo <- function() {
  # Define the tree from a Newick format string
  myTree <- ape::read.tree(
    text = "(A:0.1,B:0.2,(C:0.3,D:0.4)E:0.5);"
  )
  
  # Plot the tree with tip labels
  plot(myTree, show.tip.label = TRUE, show.node.label = TRUE)
  
  # Add edge labels showing rounded branch lengths
  edgelabels(round(myTree$edge.length, 1), cex=1)  # Adjust cex for size of text
}

# calling our function
NewickDemo()

# Next, we write a function that will generate the Newick format of a Hemoglobin subunit alpha tree
HBATree <- function(filename = "Output/HBA1/HBA1.png", width = 4000, height = 3000) {
  
  # Start PNG device
  png(filename, width = width, height = height)
  
  # Read the tree from a pre-specified file
  tree <- read.tree("Output/HBA1/hba1.tre")
  
  # Plot the tree
  plot(tree, edge.color = "black", show.node.label = FALSE, show.tip.label = TRUE, cex=2)
  
  # Add edge labels showing rounded branch lengths
  edgelabels(round(tree$edge.length, 1), cex=1)  # Adjust cex for size of text
  
  # Close the PNG device
  dev.off()
}

# Calling our function
HBATree()


# We produce a plot of SARS-CoV-2 genomes from the UK, colored by year
COVIDByYear <- function(filename = "Output/UK-SARS-CoV-2/COVIDByYear.png", width = 6000, height = 4500) {

  # Start PNG device
  png(filename, width = width, height = height)
  
  # Read the tree from a pre-specified file
  tree <- read.tree("Output/UK-SARS-CoV-2/sars-cov-2.tre")

  # Define a color palette for different years
  color_palette <- c("2020" = "green", "2021" = "blue", "2022" = "purple", "2023" = "red", "2024" = "orange")
  
  # Extract years from the tip labels of the tree
  years <- substr(tree$tip.label, 1, 4)
  
  # Assign colors based on the extracted years
  node_colors <- color_palette[years]
  
  # Plot the tree
  plot(tree, edge.color = "black", show.node.label = TRUE, show.tip.label = FALSE)
  
  # Add colored node labels
  tiplabels(pch = 19, col = node_colors, cex = 2, srt = 45, adj = c(1, 0.5), font = 2)
  
  # Add a legend to the plot
  legend("topleft",               # Position of the legend
         legend = names(color_palette),  # The text in the legend, pulling from the names of the color_palette vector
         col = color_palette,      # The colors for the symbols in the legend
         pch = 19,                 # Type of point to use, same as in nodelabels
         title = "Year",           # Title of the legend
         cex = 8)                # Font size of the text in the legend
  
  # Close the PNG device
  dev.off()
}

#calling our function
COVIDByYear()


