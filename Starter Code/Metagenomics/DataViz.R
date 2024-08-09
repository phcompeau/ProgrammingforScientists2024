# This R script takes a distance matrix between genomics samples as input.
# It produces two plots. First, a heat map of the distances in the samples. Second, a PCoA plot
# colored by the season.

# Import needed libraries. Please install these in R beforehand using install.packages("package_name")

library(ggcorrplot)
library(reshape)
library(stringr)
library(ggplot2)
library(ape) #ape library to compute PCoA of our matrix

# Now set working directory. This should be wherever the files are stored and is the only line that the user needs to edit.

# EDIT THIS LINE
setwd("/Users/phillipcompeau/go/src/Metagenomics_Final")

# Part 1: generate box plots of evenness by season using Simpson's index.

# first, read in the table

simpsonTable <- read.csv(file="Matrices/SimpsonMatrix.csv")

simpsonColumns <- data.frame(simpsonTable)

# Add column to our data frame to represent the season
Season <- sub("\\_.*", "", simpsonColumns$Sample) # parse out just the season from sample name
cbind(simpsonColumns, Season) # adding column

# Now we plot box plots where we have evenness lumped by season.
ggplot(simpsonColumns, aes(x=Season, y=SimpsonsIndex, fill=Season)) + geom_boxplot() +ggtitle("Simpson's Index over season")
ggsave("Plots/SimpsonsBoxPlots.png")


# Part 2: generate a heat map of the distance values

# Read in the file and process the table.
table <- read.csv(file="Matrices/BetaDiversityMatrix.csv")

# trim the first column out because it only contains names 
table <- table[-c(1)]

matrix <- as.matrix(table)

# the following code is just all the technical stuff to build a heatmap out of the distance matrix.
co=melt(matrix)
head(co)
ggplot(co, aes(X1, X2)) + # x and y axes => Var1 and Var2
geom_tile(aes(fill = value)) + # background colours are mapped according to the value column
scale_fill_gradient2(low = "#6D9EC1", 
                     mid = "white", 
                     high = "#E46726", 
                     midpoint = 0.5, limit= c(0,1.0)) + 
theme(panel.grid.major.x=element_blank(), #no gridlines
      panel.grid.minor.x=element_blank(), 
      panel.grid.major.y=element_blank(), 
      panel.grid.minor.y=element_blank(),
      panel.background=element_rect(fill="white"), # background=white
      axis.text.x = element_text(angle=90, hjust = 1,vjust=1,size = 8,face = "bold"),
      plot.title = element_text(size=20,face="bold"),
      axis.text.y = element_text(size = 8,face = "bold")) + 
ggtitle("Distance Heat Map") + 
theme(legend.title=element_text(face="bold", size=14)) + 
scale_x_discrete(name="") +
scale_y_discrete(name="")
ggsave("Plots/HeatMap.png")

# Part 3: generate a PCoA plot of the data and color-code by season.

# Read in the file and process the table.
table2 <- read.csv(file="Matrices/BetaDiversityMatrix.csv")

#trim the first column out because it only contains names 
table2 <- table2[-c(1)]
table2 <- table2[, -50] # trim out weird extra column at end of the matrix file

matrix2 <- as.matrix(table2)

pcoa_data <- pcoa(matrix2, correction="none", rn=NULL)
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

colnames(table2)

# Add column to our data frame to represent the season
Season2 <- sub("\\_.*", "", colnames(table2)) # parse out just the season from sample name
cbind(pcoa_vectors, Season2) # adding column

# Now, plot the data, colored by season.
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2, color=Season2)) + geom_point()
ggsave("Plots/PCoA.png")
