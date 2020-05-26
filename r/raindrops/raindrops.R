raindrops <- function(number) {
  r <- ""
  divisors <- c(3, 5, 7)
  sounds <- c("Pling", "Plang", "Plong")
  for (i in 1:length(divisors)) {
    if (number %% divisors[i] == 0) {
      r <- paste0(r, sounds[i])
    }
  }
  if (r == "") {
    return(paste0(number))
  }
  return(r)
}
