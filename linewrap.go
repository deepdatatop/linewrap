package linewrap

import(
	"strings"
)

func charType( ch rune ) (ty int) {//1: multibyte 2: digit 3: letter 4: space 5: return 6: newline 0: other
	ty = 0
	if len(string(ch))>1 {
		ty = 1
	}else if ch>=48 && ch<=57 {
		ty = 2
	}else if (ch>=65 && ch<=90) || (ch>=97 && ch<=122) {
		ty = 3
	}else if ch=='\t' || ch==' ' {
		ty = 4
	}else if ch=='\r' {
		ty = 5
	}else if ch=='\n' {
		ty = 6
	}
   return
}

func strNextWord( s []rune,i,n int ) (j int,curword string){
	j = -1
	if i<n {
		ch := s[i]
		ty := charType( ch )
		j = i+1
		if ty>=2 {
			for j<n {
				if ty<5 && charType( s[j] )==ty {
					j++
				}else if (ty==2 || ty==3) && s[j]=='-' {
					j++
				}else{
					break
				}
			}
		}
		curword = string(s[i:j])
	}
	return
}

/*func measureTextWidth(fontname,text string,fontsize int) (w float64){		//parameter func sample
	canvas.SetFont(fontname, "", fontsize)
	w = canvas.MeasureTextWidth( text )
	return
}*/

func Split2MultiLine( line,fontname string, width float64,fontsize int, measureTextWidth func(txt,fontname string,fontsize int) float64 ) (lns []string) {
	ln := ""
	s := []rune( line )
	i,n := 0,len( s )
	for i<n {
		j,wd := strNextWord( s,i,n )
		if j>0 {
			wd = strings.Replace(wd,"\t","　　",-1)
			if wd=="\r" {
			}else if wd=="\n" {
				if len(ln)>0 {
					lns = append( lns,ln )
					ln = ""
				}
			}else{
				if len(wd)>0 {
					advance := measureTextWidth( ln+wd,fontname,fontsize )	//func
					if advance<=width || len(ln)==0 {
						ln += wd
					}else{
						lns = append( lns,ln )
						ln = wd
					}
				}
			}
			i = j	
		}else{
			break
		}
	}
	lns = append( lns,ln )
	return
}