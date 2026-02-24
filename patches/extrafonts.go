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
				.locals 8
				{{range .}}
				# 开始加载字体: {{.Name}}
				const-string v1, "{{.Name}}"
				
				{{if .Regular -}}
				# 构建Regular字体路径
				new-instance v2, Ljava/lang/StringBuilder;
				invoke-direct {v2}, Ljava/lang/StringBuilder;-><init>()V
				invoke-static {}, Landroid/os/Environment;->getExternalStorageDirectory()Ljava/io/File;
				move-result-object v7
				invoke-virtual {v7}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;
				move-result-object v7
				invoke-virtual {v2, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v2
				const-string v7, "/Lithium/fonts/{{.Base}}-Regular.ttf"
				invoke-virtual {v2, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v2
				invoke-virtual {v2}, Ljava/lang/StringBuilder;->toString()Ljava/lang/String;
				move-result-object v2
				
				# 检查文件是否存在
				new-instance v7, Ljava/io/File;
				invoke-direct {v7, v2}, Ljava/io/File;-><init>(Ljava/lang/String;)V
				invoke-virtual {v7}, Ljava/io/File;->exists()Z
				move-result v7
				if-eqz v7, :cond_regular_ok
				const/4 v2, 0x0
				:cond_regular_ok
				{{- else -}}
				const/4 v2, 0x0
				{{- end}}
				
				{{if .Bold -}}
				new-instance v3, Ljava/lang/StringBuilder;
				invoke-direct {v3}, Ljava/lang/StringBuilder;-><init>()V
				invoke-static {}, Landroid/os/Environment;->getExternalStorageDirectory()Ljava/io/File;
				move-result-object v7
				invoke-virtual {v7}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;
				move-result-object v7
				invoke-virtual {v3, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v3
				const-string v7, "/Lithium/fonts/{{.Base}}-Bold.ttf"
				invoke-virtual {v3, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v3
				invoke-virtual {v3}, Ljava/lang/StringBuilder;->toString()Ljava/lang/String;
				move-result-object v3
				
				new-instance v7, Ljava/io/File;
				invoke-direct {v7, v3}, Ljava/io/File;-><init>(Ljava/lang/String;)V
				invoke-virtual {v7}, Ljava/io/File;->exists()Z
				move-result v7
				if-eqz v7, :cond_bold_ok
				const/4 v3, 0x0
				:cond_bold_ok
				{{- else -}}
				const/4 v3, 0x0
				{{- end}}
				
				{{if .Italic -}}
				new-instance v4, Ljava/lang/StringBuilder;
				invoke-direct {v4}, Ljava/lang/StringBuilder;-><init>()V
				invoke-static {}, Landroid/os/Environment;->getExternalStorageDirectory()Ljava/io/File;
				move-result-object v7
				invoke-virtual {v7}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;
				move-result-object v7
				invoke-virtual {v4, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v4
				const-string v7, "/Lithium/fonts/{{.Base}}-Italic.ttf"
				invoke-virtual {v4, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v4
				invoke-virtual {v4}, Ljava/lang/StringBuilder;->toString()Ljava/lang/String;
				move-result-object v4
				
				new-instance v7, Ljava/io/File;
				invoke-direct {v7, v4}, Ljava/io/File;-><init>(Ljava/lang/String;)V
				invoke-virtual {v7}, Ljava/io/File;->exists()Z
				move-result v7
				if-eqz v7, :cond_italic_ok
				const/4 v4, 0x0
				:cond_italic_ok
				{{- else -}}
				const/4 v4, 0x0
				{{- end}}
				
				{{if .BoldItalic -}}
				new-instance v5, Ljava/lang/StringBuilder;
				invoke-direct {v5}, Ljava/lang/StringBuilder;-><init>()V
				invoke-static {}, Landroid/os/Environment;->getExternalStorageDirectory()Ljava/io/File;
				move-result-object v7
				invoke-virtual {v7}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;
				move-result-object v7
				invoke-virtual {v5, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v5
				const-string v7, "/Lithium/fonts/{{.Base}}-BoldItalic.ttf"
				invoke-virtual {v5, v7}, Ljava/lang/StringBuilder;->append(Ljava/lang/String;)Ljava/lang/StringBuilder;
				move-result-object v5
				invoke-virtual {v5}, Ljava/lang/StringBuilder;->toString()Ljava/lang/String;
				move-result-object v5
				
				new-instance v7, Ljava/io/File;
				invoke-direct {v7, v5}, Ljava/io/File;-><init>(Ljava/lang/String;)V
				invoke-virtual {v7}, Ljava/io/File;->exists()Z
				move-result v7
				if-eqz v7, :cond_bolditalic_ok
				const/4 v5, 0x0
				:cond_bolditalic_ok
				{{- else -}}
				const/4 v5, 0x0
				{{- end}}
				
				const/16 v6, {{.Script.Flags | printf "%#x"}} # {{.Script}}

				# 创建Font对象
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
					.locals 6

					# 添加异常处理
					:try_start_0
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
					const/16 v4, {{$x.Script.Flags | printf "%#x"}} # {{$x.Script}}
					if-nez v0, :filter
					{{- end}}
					{{- end}}

					const/16 v4, 0x0

					:filter
					new-instance v0, Ljava/util/ArrayList;
					invoke-direct {v0}, Ljava/util/ArrayList;-><init>()V

					invoke-static {}, Lcom/faultexception/reader/fonts/Fonts;->getFonts()Ljava/util/List;
					move-result-object v1
					
					if-nez v1, :cond_empty
					return-object v0
					
					:cond_empty
					invoke-interface {v1}, Ljava/util/List;->iterator()Ljava/util/Iterator;
					move-result-object v1

					:filter_next
					invoke-interface {v1}, Ljava/util/Iterator;->hasNext()Z
					move-result v2
					if-eqz v2, :filter_done

					invoke-interface {v1}, Ljava/util/Iterator;->next()Ljava/lang/Object;
					move-result-object v2
					check-cast v2, Lcom/faultexception/reader/fonts/Font;
					
					if-nez v2, :filter_next

					iget v3, v2, Lcom/faultexception/reader/fonts/Font;->scripts:I
					and-int v5, v3, v4
					if-eq v5, v4, :cond_add
					goto :filter_next
					
					:cond_add
					invoke-interface {v0, v2}, Ljava/util/List;->add(Ljava/lang/Object;)Z
					goto :filter_next
				
					:filter_done
					return-object v0
					:try_end_0
					.catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

					:catch_0
					move-exception v0
					new-instance v1, Ljava/util/ArrayList;
					invoke-direct {v1}, Ljava/util/ArrayList;-><init>()V
					return-object v1
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
