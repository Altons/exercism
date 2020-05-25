raindrops <- function(number) {
  r<-""
  if( number%%3 == 0) {
    r <-paste0(r,"Pling")
  }
  if( number%%5 == 0) {
    r <-paste0(r,"Plang")
  }
  if( number%%7 == 0) {
    r <-paste0(r,"Plong")
  }
  if (r == "") {
    return(paste0(number))
  }
  return(r)
}
