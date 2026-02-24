// # Extra fonts
//
// Add additional fonts.
package patches

import (
	"io"

	"github.com/pgaskin/lithiumpatch/fonts"
	. "github.com/pgaskin/lithiumpatch/patches/patchdef"
)

func init() {
	Register("extrafonts", extrafonts{})
}

type extrafonts struct{}

func (extrafonts) Do(apk string, diffwriter io.Writer) error {
	xs := fonts.All()
	if len(xs) == 0 {
		return nil
	}
	var pt []Instruction
	
	pt = append(pt, PatchFile("smali/com/faultexception/reader/fonts/Fonts.smali",
		InMethod("<clinit>()V",
			ReplaceStringAppend(
				FixIndent("\n"+`
					invoke-static {v0}, Ljava/util/Arrays;->asList([Ljava/lang/Object;)Ljava/util/List;

					move-result-object v0
				`),
				FixIndent("\n"+`
					new-instance v1, Ljava/util/ArrayList;
					invoke-direct {v1, v0}, Ljava/util/ArrayList;-><init>(Ljava/util/Collection;)V
					move-object v0, v1

					invoke-static {v0}, Lcom/faultexception/reader/fonts/Fonts;->initCustomFonts(Ljava/util/ArrayList;)V
				`),
			),
		),
		AppendString(
			FixIndent(ExecuteTemplate("\n"+`
			.method private static initCustomFonts(Ljava/util/ArrayList;)V
				.locals 7
				{{range .}}
				const-string v1, "{{.Name}}"
				
				{{if .Regular -}}
				const-string v2, "/sdcard/Lithium/fonts/{{.Base}}-Regular.ttf"
				{{- else -}}
				const/4 v2, 0x0
				{{- end}}
				
				{{if .Bold -}}
				const-string v3, "/sdcard/Lithium/fonts/{{.Base}}-Bold.ttf"
				{{- else -}}
				const/4 v3, 0x0
				{{- end}}
				
				{{if .Italic -}}
				const-string v4, "/sdcard/Lithium/fonts/{{.Base}}-Italic.ttf"
				{{- else -}}
				const/4 v4, 0x0
				{{- end}}
				
				{{if .BoldItalic -}}
				const-string v5, "/sdcard/Lithium/fonts/{{.Base}}-BoldItalic.ttf"
				{{- else -}}
				const/4 v5, 0x0
				{{- end}}
				
				const/16 v6, {{.Script.Flags | printf "%#x"}}

				new-instance v0, Lcom/faultexception/reader/fonts/Font;
				invoke-direct/range {v0 .. v6}, Lcom/faultexception/reader/fonts/Font;-><init>(Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;I)V
				invoke-virtual {p0, v0}, Ljava/util/ArrayList;->add(Ljava/lang/Object;)Z
				{{end}}
				return-void
			.end method
			`, xs)),
		),
		InMethod("getCompatibleFonts(Ljava/lang/String;)Ljava/util/List;",
			ReplaceWith(
				FixIndent(ExecuteTemplate("\n"+`
					.locals 5

					const-string v0, "-"
					invoke-virtual {p0, v0}, Ljava/lang/String;->split(Ljava/lang/String;)[Ljava/lang/String;
					move-result-object p0

					const/4 v0, 0x0
					aget-object p0, p0, v0
					invoke-virtual {p0}, Ljava/lang/String;->toLowerCase()Ljava/lang/String;
					move-result-object p0

					{{range $x := .}}
					{{range $x.Language}}
					const-string v0, "{{.}}"
					invoke-virtual {p0, v0}, Ljava/lang/String;->equals(Ljava/lang/Object;)Z
					move-result v0
					if-eqz v0, :script_match
					{{- end}}
					{{- end}}
					
					const/16 v4, 0x0
					goto :filter_start

					{{range $x := .}}
					:script_match_{{$x.Script}}
					const/16 v4, {{$x.Script.Flags | printf "%#x"}}
					{{end}}

					:filter_start
					new-instance v0, Ljava/util/ArrayList;
					invoke-direct {v0}, Ljava/util/ArrayList;-><init>()V

					invoke-static {}, Lcom/faultexception/reader/fonts/Fonts;->getFonts()Ljava/util/List;
					move-result-object v1
					invoke-interface {v1}, Ljava/util/List;->iterator()Ljava/util/Iterator;
					move-result-object v1

					:cond_loop
					invoke-interface {v1}, Ljava/util/Iterator;->hasNext()Z
					move-result v2
					if-nez v2, :return_result

					invoke-interface {v1}, Ljava/util/Iterator;->next()Ljava/lang/Object;
					move-result-object v2
					check-cast v2, Lcom/faultexception/reader/fonts/Font;

					iget v3, v2, Lcom/faultexception/reader/fonts/Font;->scripts:I
					and-int/2addr v3, v4
					if-eq v3, v4, :cond_loop

					invoke-interface {v0, v2}, Ljava/util/List;->add(Ljava/lang/Object;)Z
					goto :cond_loop

					:return_result
					return-object v0
				`, []struct {
					Script   fonts.Script
					Language []string
				}{
					{fonts.FontScriptLatin, []string{"eng", "en"}},
					{fonts.FontScriptCyrillic, []string{"rus", "ru"}},
					{fonts.FontScriptGreek, []string{"gre", "ell", "el"}},
					{fonts.FontScriptThai, []string{"tha", "th"}},
				})),
			),
		),
	))
	
	for _, x := range pt {
		if err := x.Do(apk, diffwriter); err != nil {
			return err
		}
	}
	return nil
}
