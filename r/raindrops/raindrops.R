raindrops <- function(number) {
  divisors <- c(3, 5, 7)
  sounds <- c("Pling", "Plang", "Plong")
  r <- paste(sounds[number %% divisors ==0],collapse="")
  if (r == "") {
    return(paste0(number))
  }
  return(r)
}
