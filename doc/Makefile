# Makefile for Sphinx documentation
#

# You can set these variables from the command line.
SPHINXOPTS    =
SPHINXBUILD   = sphinx-build
PAPER         =
BUILDDIR      = build

# Internal variables.
PAPEROPT_a4     = -D latex_paper_size=a4
PAPEROPT_letter = -D latex_paper_size=letter
ALLSPHINXOPTS   = -d $(BUILDDIR)/doctrees $(PAPEROPT_$(PAPER)) $(SPHINXOPTS) source

.PHONY: help clean html dirhtml singlehtml pickle json htmlhelp qthelp devhelp epub latex latexpdf text man changes linkcheck doctest

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  html       to make standalone HTML files"
	@echo "  dirhtml    to make HTML files named index.html in directories"
	@echo "  offhtml    to make standalone HTML files without fetching the new theme files"
	@echo "  singlehtml to make a single large HTML file"
	@echo "  pickle     to make pickle files"
	@echo "  json       to make JSON files"
	@echo "  htmlhelp   to make HTML files and a HTML help project"
	@echo "  qthelp     to make HTML files and a qthelp project"
	@echo "  devhelp    to make HTML files and a Devhelp project"
	@echo "  epub       to make an epub"
	@echo "  latex      to make LaTeX files, you can set PAPER=a4 or PAPER=letter"
	@echo "  latexpdf   to make LaTeX files and run them through pdflatex"
	@echo "  text       to make text files"
	@echo "  man        to make manual pages"
	@echo "  changes    to make an overview of all changed/added/deprecated items"
	@echo "  linkcheck  to check all external links for integrity"
	@echo "  doctest    to run all doctests embedded in the documentation (if enabled)"

clean:
	-rm -rf $(BUILDDIR)/*
	-rm -f source/percona-theme/built

source/percona-theme/built:
	@echo "Downloading percona-theme ..."
	@wget -O percona-theme.tar.gz https://www.percona.com/docs/theme-1-4/percona-xtradb-cluster/5.6/
	@echo "Extracting theme."
	@tar -zxf percona-theme.tar.gz
	@rm -rf source/percona-theme
	@mv percona-theme-1-4 source/percona-theme
	@rm percona-theme.tar.gz
	@sed -i 's^{{ toc }}^<style>.select-wrapper { display: inline-flex; flex-direction: column; } #custom_select { margin-bottom: 5px; } #custom_select_list.select-hidden { display: none;} #custom_select_list { display: inline-flex; flex-direction: column; padding-left: 0; } .custom-select__option:not(:last-child) { margin-bottom: 5x; } </style><section class="select-wrapper"><div class="btn btn-primary" id="custom_select"><span class="select-active-text"></span><span class="glyphicon glyphicon-chevron-down"></span></div><ul class="select-hidden" id="custom_select_list"><li><a class="custom-select__option"href="javascript:window.pmmDocRedirect()">PMM 1</a></li><li><a class="custom-select__option"  href="javascript:window.pmm2DocRedirect()">PMM2</a></li></ul></section><script>window.pmm2DocRedirect = () => { const docLink = document.querySelector("#another-doc-version-link"); let newLink =  "https://www.percona.com/doc/percona-monitoring-and-management/2.x/index.html"; if (docLink) { if (docLink.href) { newLink = docLink.href; } } window.location.href = newLink;};window.pmmDocRedirect = () => { const docLink = document.querySelector("#another-doc-version-link"); let newLink = "https://www.percona.com/doc/percona-monitoring-and-management/index.html"; if (docLink) { if (docLink.dataset.location) {            newLink = docLink.dataset.location; } } window.location.href = newLink;};</script><script>setTimeout(() => { makeSelect(); }, 500); function makeSelect() { const custom_select = document.getElementById("custom_select"); const select_active_option = custom_select.getElementsByClassName("select-active-text")[0]; const custom_select_list = document.getElementById("custom_select_list"); select_active_option.innerHTML = window.location.href.includes("2.x") ?        custom_select_list.getElementsByClassName("custom-select__option")[1].innerHTML : custom_select_list.getElementsByClassName("custom-select__option")[0].innerHTML; document.addEventListener("click", event => {        if (event.target.parentElement.id === "custom_select" || event.target.id === "custom_select") { custom_select_list.classList.toggle("select-hidden") } if (Array.from(event.target.classList).includes("custom-select__option")) { select_active_option.innerHTML = event.target.innerHTML; } if (event.target.id !== "custom_select") { if (event.target.parentElement.id !== "custom_select") { custom_select_list.classList.add("select-hidden") } } });}</script> {{ toc }}^' source/percona-theme/localtoc.html
	@touch source/percona-theme/built

theme:  source/percona-theme/built

html:   theme
	@echo "Building html doc"
	$(SPHINXBUILD) -b html $(ALLSPHINXOPTS) $(BUILDDIR)/html
	@echo
	@echo "Build finished. The HTML pages are in $(BUILDDIR)/html."

dirhtml:
	$(SPHINXBUILD) -b dirhtml $(ALLSPHINXOPTS) $(BUILDDIR)/dirhtml
	@echo
	@echo "Build finished. The HTML pages are in $(BUILDDIR)/dirhtml."

offhtml:
	$(SPHINXBUILD) -b html $(ALLSPHINXOPTS) $(BUILDDIR)/html
	@echo
	@echo "Build finished. The HTML pages are in $(BUILDDIR)/html."

singlehtml:
	$(SPHINXBUILD) -b singlehtml $(ALLSPHINXOPTS) $(BUILDDIR)/singlehtml
	@echo
	@echo "Build finished. The HTML page is in $(BUILDDIR)/singlehtml."

pickle:
	$(SPHINXBUILD) -b pickle $(ALLSPHINXOPTS) $(BUILDDIR)/pickle
	@echo
	@echo "Build finished; now you can process the pickle files."

json:
	$(SPHINXBUILD) -b json $(ALLSPHINXOPTS) $(BUILDDIR)/json
	@echo
	@echo "Build finished; now you can process the JSON files."

htmlhelp:
	$(SPHINXBUILD) -b htmlhelp $(ALLSPHINXOPTS) $(BUILDDIR)/htmlhelp
	@echo
	@echo "Build finished; now you can run HTML Help Workshop with the" \
	      ".hhp project file in $(BUILDDIR)/htmlhelp."

qthelp:
	$(SPHINXBUILD) -b qthelp $(ALLSPHINXOPTS) $(BUILDDIR)/qthelp
	@echo
	@echo "Build finished; now you can run "qcollectiongenerator" with the" \
	      ".qhcp project file in $(BUILDDIR)/qthelp, like this:"
	@echo "# qcollectiongenerator $(BUILDDIR)/qthelp/PerconaServer.qhcp"
	@echo "To view the help file:"
	@echo "# assistant -collectionFile $(BUILDDIR)/qthelp/PerconaServer.qhc"

devhelp:
	$(SPHINXBUILD) -b devhelp $(ALLSPHINXOPTS) $(BUILDDIR)/devhelp
	@echo
	@echo "Build finished."
	@echo "To view the help file:"
	@echo "# mkdir -p $$HOME/.local/share/devhelp/PerconaServer"
	@echo "# ln -s $(BUILDDIR)/devhelp $$HOME/.local/share/devhelp/PerconaServer"
	@echo "# devhelp"

epub:
	$(SPHINXBUILD) -b epub $(ALLSPHINXOPTS) $(BUILDDIR)/epub
	@echo
	@echo "Build finished. The epub file is in $(BUILDDIR)/epub."

latex:
	$(SPHINXBUILD) -b latex $(ALLSPHINXOPTS) $(BUILDDIR)/latex
	@echo
	@echo "Build finished; the LaTeX files are in $(BUILDDIR)/latex."
	@echo "Run \`make' in that directory to run these through (pdf)latex" \
	      "(use \`make latexpdf' here to do that automatically)."

latexpdf:
	$(SPHINXBUILD) -b latex $(ALLSPHINXOPTS) $(BUILDDIR)/latex
	@echo "Running LaTeX files through pdflatex..."
	make -C $(BUILDDIR)/latex all-pdf
	@echo "pdflatex finished; the PDF files are in $(BUILDDIR)/latex."

text:
	$(SPHINXBUILD) -b text $(ALLSPHINXOPTS) $(BUILDDIR)/text
	@echo
	@echo "Build finished. The text files are in $(BUILDDIR)/text."

man:
	$(SPHINXBUILD) -b man $(ALLSPHINXOPTS) $(BUILDDIR)/man
	@echo
	@echo "Build finished. The manual pages are in $(BUILDDIR)/man."

changes:
	$(SPHINXBUILD) -b changes $(ALLSPHINXOPTS) $(BUILDDIR)/changes
	@echo
	@echo "The overview file is in $(BUILDDIR)/changes."

linkcheck:
	$(SPHINXBUILD) -b linkcheck $(ALLSPHINXOPTS) $(BUILDDIR)/linkcheck
	@echo
	@echo "Link check complete; look for any errors in the above output " \
	      "or in $(BUILDDIR)/linkcheck/output.txt."

doctest:
	$(SPHINXBUILD) -b doctest $(ALLSPHINXOPTS) $(BUILDDIR)/doctest
	@echo "Testing of doctests in the sources finished, look at the " \
	      "results in $(BUILDDIR)/doctest/output.txt."
