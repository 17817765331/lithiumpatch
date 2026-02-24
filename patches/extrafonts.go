// # Extra fonts
//
// Add additional fonts and support external font directories.
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
	for _, f := range xs {
		if f.Regular != nil {
			pt = append(pt, WriteFile("assets/fonts/"+f.Base+"-Regular.ttf", f.Regular))
		}
		if f.Bold != nil {
			pt = append(pt, WriteFile("assets/fonts/"+f.Base+"-Bold.ttf", f.Bold))
		}
		if f.Italic != nil {
			pt = append(pt, WriteFile("assets/fonts/"+f.Base+"-Italic.ttf", f.Italic))
		}
		if f.BoldItalic != nil {
			pt = append(pt, WriteFile("assets/fonts/"+f.Base+"-BoldItalic.ttf", f.BoldItalic))
		}
	}
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
				.locals 8
				.param p0, "fonts"    # Ljava/util/ArrayList;
				{{range .}}
				const-string v1, "{{.Name}}"
				{{if .Regular -}}
				const-string v2, "{{.Base}}-Regular.ttf"
				{{- else -}}
				const/4 v2, 0x0
				{{- end}}
				{{if .Bold -}}
				const-string v3, "{{.Base}}-Bold.ttf"
				{{- else -}}
				const/4 v3, 0x0
				{{- end}}
				{{if .Italic -}}
				const-string v4, "{{.Base}}-Italic.ttf"
				{{- else -}}
				const/4 v4, 0x0
				{{- end}}
				{{if .BoldItalic -}}
				const-string v5, "{{.Base}}-BoldItalic.ttf"
				{{- else -}}
				const/4 v5, 0x0
				{{- end}}
				const/16 v6, {{.Script.Flags | printf "%#x"}} # {{.Script}}

				new-instance v0, Lcom/faultexception/reader/fonts/Font;
				invoke-direct/range {v0 .. v6}, Lcom/faultexception/reader/fonts/Font;-><init>(Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;I)V
				invoke-virtual {p0, v0}, Ljava/util/ArrayList;->add(Ljava/lang/Object;)Z
				{{end}}

				# Load external fonts from /sdcard/Lithium/fonts/
				:try_start_0
				new-instance v7, Ljava/io/File;
				invoke-static {}, Landroid/os/Environment;->getExternalStorageDirectory()Ljava/io/File;
				move-result-object v0
				const-string v1, "Lithium/fonts"
				invoke-direct {v7, v0, v1}, Ljava/io/File;-><init>(Ljava/io/File;Ljava/lang/String;)V
				invoke-static {v7}, Lcom/faultexception/reader/fonts/Fonts;->scanExternalFonts(Ljava/io/File;)V
				:try_end_0
				.catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

				:goto_0
				# Load external fonts from app-specific directory
				:try_start_1
				invoke-static {}, Lcom/faultexception/reader/MainApplication;->getInstance()Lcom/faultexception/reader/MainApplication;
				move-result-object v7
				const/4 v0, 0x0
				invoke-virtual {v7, v0}, Landroid/content/Context;->getExternalFilesDir(Ljava/lang/String;)Ljava/io/File;
				move-result-object v7

				if-eqz v7, :cond_0
				invoke-static {v7}, Lcom/faultexception/reader/fonts/Fonts;->scanExternalFonts(Ljava/io/File;)V
				:cond_0
				:try_end_1
				.catch Ljava/lang/Exception; {:try_start_1 .. :try_end_1} :catch_1

				:goto_1
				return-void

				:catch_0
				move-exception v7
				goto :goto_0

				:catch_1
				move-exception v7
				goto :goto_1
			.end method

			.method private static loadExternalFont(Ljava/io/File;)V
			    .locals 5
			    .param p0, "fontFile"    # Ljava/io/File;

			    .prologue
			    invoke-virtual {p0}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;
			    move-result-object v0

			    :try_start_0
			    invoke-static {v0}, Landroid/graphics/Typeface;->createFromFile(Ljava/lang/String;)Landroid/graphics/Typeface;
			    move-result-object v1

			    if-nez v1, :cond_0
			    return-void

			    :cond_0
			    invoke-virtual {p0}, Ljava/io/File;->getName()Ljava/lang/String;
			    move-result-object v2
			    const/16 v3, 0x2e
			    invoke-virtual {v2, v3}, Ljava/lang/String;->lastIndexOf(I)I
			    move-result v3
			    if-lez v3, :cond_1
			    const/4 v4, 0x0
			    invoke-virtual {v2, v4, v3}, Ljava/lang/String;->substring(II)Ljava/lang/String;
			    move-result-object v2

			    :cond_1
			    new-instance v3, Lcom/faultexception/reader/fonts/Font;
			    const/4 v4, 0x0
			    invoke-direct {v3, v2, v0, v4, v4}, Lcom/faultexception/reader/fonts/Font;-><init>(Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;I)V

			    invoke-static {}, Lcom/faultexception/reader/fonts/Fonts;->getFonts()Ljava/util/List;
			    move-result-object v4
			    invoke-interface {v4, v3}, Ljava/util/List;->add(Ljava/lang/Object;)Z
			    :try_end_0
			    .catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

			    :goto_0
			    return-void

			    :catch_0
			    move-exception v1
			    goto :goto_0
			.end method

			.method private static scanExternalFonts(Ljava/io/File;)V
			    .locals 7
			    .param p0, "fontDir"    # Ljava/io/File;

			    .prologue
			    :try_start_0
			    invoke-virtual {p0}, Ljava/io/File;->exists()Z
			    move-result v0
			    if-nez v0, :cond_0
			    return-void

			    :cond_0
			    invoke-virtual {p0}, Ljava/io/File;->isDirectory()Z
			    move-result v0
			    if-nez v0, :cond_1
			    return-void

			    :cond_1
			    invoke-virtual {p0}, Ljava/io/File;->listFiles()[Ljava/io/File;
			    move-result-object v1

			    if-nez v1, :cond_2
			    return-void

			    :cond_2
			    array-length v2, v1
			    const/4 v3, 0x0

			    :goto_0
			    if-ge v3, v2, :cond_4
			    aget-object v4, v1, v3

			    invoke-virtual {v4}, Ljava/io/File;->isDirectory()Z
			    move-result v5
			    if-eqz v5, :cond_3
			    invoke-static {v4}, Lcom/faultexception/reader/fonts/Fonts;->scanExternalFonts(Ljava/io/File;)V
			    goto :goto_1

			    :cond_3
			    invoke-virtual {v4}, Ljava/io/File;->getName()Ljava/lang/String;
			    move-result-object v5
			    invoke-virtual {v5}, Ljava/lang/String;->toLowerCase()Ljava/lang/String;
			    move-result-object v5
			    const-string v6, ".ttf"
			    invoke-virtual {v5, v6}, Ljava/lang/String;->endsWith(Ljava/lang/String;)Z
			    move-result v5
			    if-eqz v5, :cond_5

			    invoke-static {v4}, Lcom/faultexception/reader/fonts/Fonts;->loadExternalFont(Ljava/io/File;)V

			    :goto_1
			    add-int/lit8 v3, v3, 0x1
			    goto :goto_0

			    :cond_4
			    return-void

			    :cond_5
			    goto :goto_1
			    :try_end_0
			    .catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

			    :catch_0
			    move-exception v0
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

					const v0, 0
					aget-object p0, p0, v0
					invoke-virtual {p0}, Ljava/lang/String;->toLowerCase()Ljava/lang/String;
					move-result-object p0

					{{range $x := .}}
					{{range $x.Language}}
					const-string v0, "{{.}}"
					invoke-virtual {p0, v0}, Ljava/lang/String;->equals(Ljava/lang/Object;)Z
					move-result v0
					const/16 v4, {{$x.Script.Flags | printf "%#x"}} # {{$x.Script}}
					if-nez v0, :filter
					{{- end}}
					{{- end}}

					const/16 v4, 0 # any (default)

					:filter
					new-instance v0, Ljava/util/ArrayList;
					invoke-direct {v0}, Ljava/util/ArrayList;-><init>()V

					invoke-static {}, Lcom/faultexception/reader/fonts/Fonts;->getFonts()Ljava/util/List;
					move-result-object v1
					invoke-interface {v1}, Ljava/util/List;->iterator()Ljava/util/Iterator;
					move-result-object v1

					:filter_next
					invoke-interface {v1}, Ljava/util/Iterator;->hasNext()Z
					move-result v2
					if-eqz v2, :filter_done

					invoke-interface {v1}, Ljava/util/Iterator;->next()Ljava/lang/Object;
					move-result-object v2
					check-cast v2, Lcom/faultexception/reader/fonts/Font;

					iget v3, v2, Lcom/faultexception/reader/fonts/Font;->scripts:I
					and-int/2addr v3, v4
					if-ne v3, v4, :filter_next

					invoke-interface {v0, v2}, Ljava/util/List;->add(Ljava/lang/Object;)Z
					goto :filter_next
				
					:filter_done
					return-object v0
				`, []struct {
					Script   fonts.Script
					Language []string
				}{
					{fonts.FontScriptLatin, []string{"eng", "en"}},
					{fonts.FontScriptCyrillic, []string{"rus", "ru"}},
					{fonts.FontScriptGreek, []string{"gre", "ell", "el"}},
					{fonts.FontScriptThai, []string{"tha", "th"}},
					// default is any
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
