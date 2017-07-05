# linewrap
string line wrap by width

Sample measure text width function

func measureTextWidth(fontname,text string,fontsize int) (w float64){
	Canvas.SetFont(fontname, fontsize)
	w = Canvas.MeasureTextWidth( text )
	return
}

Canvas can replace by image canvas, PDF canvas , Printer canvas ...etc
Canvas.MeasureTextWidth maybe  font.MeasureString(...)  ... etc

func Split2MultiLine( line,fontname string, width float64,fontsize int, 
              measureTextWidth func(txt,fontname string,fontsize int) float64 ) (lns []string)

Sample:       linewrap.Split2MultiLine( text,"fontname",width,11,measureTextWidth )
